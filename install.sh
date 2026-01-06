#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
REPO="khotcholava/zhvabu-cli"  # Update this with your GitHub username/repo
BINARY_NAME="rc"
INSTALL_DIR="/usr/local/bin"
VERSION=${1:-latest}

# Detect OS and Architecture
detect_platform() {
    OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
    ARCH="$(uname -m)"
    
    case "$ARCH" in
        x86_64)
            ARCH="amd64"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        *)
            echo -e "${RED}Unsupported architecture: $ARCH${NC}"
            exit 1
            ;;
    esac
    
    case "$OS" in
        darwin)
            OS="darwin"
            ;;
        linux)
            OS="linux"
            ;;
        *)
            echo -e "${RED}Unsupported OS: $OS${NC}"
            exit 1
            ;;
    esac
    
    echo "$OS/$ARCH"
}

# Get latest version from GitHub API
get_latest_version() {
    if [ "$VERSION" = "latest" ]; then
        VERSION_TAG=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | \
        grep '"tag_name":' | \
        sed -E 's/.*"([^"]+)".*/\1/')
        
        # Check if version tag is empty (no releases yet)
        if [ -z "$VERSION_TAG" ]; then
            echo -e "${RED}No releases found for ${REPO}${NC}" >&2
            echo -e "${YELLOW}Please create a release first, or specify a version:${NC}" >&2
            echo -e "${YELLOW}  $0 v0.1.0${NC}" >&2
            exit 1
        fi
        
        echo "$VERSION_TAG"
    else
        echo "$VERSION"
    fi
}

# Download and install binary
install_binary() {
    PLATFORM=$(detect_platform)
    OS=$(echo $PLATFORM | cut -d'/' -f1)
    ARCH=$(echo $PLATFORM | cut -d'/' -f2)
    
    VERSION_TAG=$(get_latest_version)
    
    # Validate version tag is not empty
    if [ -z "$VERSION_TAG" ]; then
        echo -e "${RED}Error: Version tag is empty${NC}"
        exit 1
    fi
    
    echo -e "${GREEN}Installing ${BINARY_NAME} ${VERSION_TAG} for ${OS}/${ARCH}...${NC}"
    
    # Determine file extension
    if [ "$OS" = "windows" ]; then
        EXT="zip"
        BINARY_FILE="${BINARY_NAME}.exe"
        ARCHIVE_NAME="${BINARY_NAME}-${OS}-${ARCH}.${EXT}"
        ALT_ARCHIVE_NAME="${BINARY_NAME}-${OS}-${ARCH}"
    else
        EXT="tar.gz"
        BINARY_FILE="${BINARY_NAME}"
        ARCHIVE_NAME="${BINARY_NAME}-${OS}-${ARCH}.${EXT}"
        ALT_ARCHIVE_NAME="${BINARY_NAME}-${OS}-${ARCH}"
    fi
    
    # Download URL (try with extension first, then without)
    DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION_TAG}/${ARCHIVE_NAME}"
    ALT_DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION_TAG}/${ALT_ARCHIVE_NAME}"
    
    # Create temp directory
    TEMP_DIR=$(mktemp -d)
    trap "rm -rf $TEMP_DIR" EXIT
    
    echo -e "${YELLOW}Downloading from ${DOWNLOAD_URL}...${NC}"
    
    # Download (try with extension first)
    DOWNLOAD_SUCCESS=false
    if curl -fsSL "$DOWNLOAD_URL" -o "${TEMP_DIR}/${ARCHIVE_NAME}" 2>/dev/null; then
        DOWNLOAD_SUCCESS=true
        DOWNLOADED_FILE="${ARCHIVE_NAME}"
    elif curl -fsSL "$ALT_DOWNLOAD_URL" -o "${TEMP_DIR}/${ALT_ARCHIVE_NAME}" 2>/dev/null; then
        DOWNLOAD_SUCCESS=true
        DOWNLOADED_FILE="${ALT_ARCHIVE_NAME}"
        echo -e "${YELLOW}Downloaded without extension, trying to extract...${NC}"
    fi
    
    if [ "$DOWNLOAD_SUCCESS" = false ]; then
        echo -e "${RED}Failed to download binary${NC}"
        echo -e "${YELLOW}Tried URLs:${NC}"
        echo -e "${YELLOW}  1. ${DOWNLOAD_URL}${NC}"
        echo -e "${YELLOW}  2. ${ALT_DOWNLOAD_URL}${NC}"
        echo -e "${YELLOW}This might mean:${NC}"
        echo -e "${YELLOW}  1. The release ${VERSION_TAG} doesn't exist yet${NC}"
        echo -e "${YELLOW}  2. The binary for ${OS}/${ARCH} wasn't uploaded${NC}"
        echo -e "${YELLOW}  3. Check available releases at: https://github.com/${REPO}/releases${NC}"
        exit 1
    fi
    
    # Extract or copy binary
    cd "$TEMP_DIR"
    
    # Check if downloaded file is already a binary or an archive
    if [ -f "${ARCHIVE_NAME}" ]; then
        # Check file type
        FILE_TYPE=$(file -b "${ARCHIVE_NAME}")
        if echo "$FILE_TYPE" | grep -qE "(gzip|Zip archive|tar)"; then
            # It's an archive, extract it
            if [ "$EXT" = "zip" ]; then
                unzip -q "${ARCHIVE_NAME}"
            else
                tar -xzf "${ARCHIVE_NAME}"
            fi
        else
            # It's already a binary, just copy it
            cp "${ARCHIVE_NAME}" "${BINARY_FILE}"
        fi
    elif [ -f "${ALT_ARCHIVE_NAME}" ]; then
        # File downloaded without extension
        FILE_TYPE=$(file -b "${ALT_ARCHIVE_NAME}")
        if echo "$FILE_TYPE" | grep -qE "(gzip|Zip archive|tar)"; then
            # It's an archive, extract it
            if [ "$EXT" = "zip" ]; then
                unzip -q "${ALT_ARCHIVE_NAME}"
            else
                tar -xzf "${ALT_ARCHIVE_NAME}"
            fi
        else
            # It's already a binary, just copy it
            cp "${ALT_ARCHIVE_NAME}" "${BINARY_FILE}"
        fi
    fi
    
    # Check if binary exists
    if [ ! -f "$BINARY_FILE" ]; then
        echo -e "${RED}Binary not found in archive${NC}"
        exit 1
    fi
    
    # Check if install directory is writable
    if [ ! -w "$(dirname $INSTALL_DIR)" ]; then
        echo -e "${YELLOW}Install directory requires sudo permissions${NC}"
        sudo mv "$BINARY_FILE" "${INSTALL_DIR}/${BINARY_NAME}"
        sudo chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
    else
        mv "$BINARY_FILE" "${INSTALL_DIR}/${BINARY_NAME}"
        chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
    fi
    
    # Verify installation
    if command -v "$BINARY_NAME" > /dev/null 2>&1; then
        INSTALLED_VERSION=$($BINARY_NAME --version 2>&1 || echo "unknown")
        echo -e "${GREEN}Successfully installed ${BINARY_NAME} ${INSTALLED_VERSION}${NC}"
        echo -e "${GREEN}Run '${BINARY_NAME} --help' to get started${NC}"
    else
        echo -e "${YELLOW}Installation completed, but ${BINARY_NAME} is not in PATH${NC}"
        echo -e "${YELLOW}Add ${INSTALL_DIR} to your PATH${NC}"
    fi
}

# Main
main() {
    echo -e "${GREEN}React CLI Installer${NC}"
    echo ""
    
    # Check for required commands
    if ! command -v curl > /dev/null 2>&1; then
        echo -e "${RED}curl is required but not installed${NC}"
        exit 1
    fi
    
    install_binary
}

main "$@"

