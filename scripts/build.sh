#!/bin/bash

# Build script for react-cli
# Builds binaries for multiple platforms

set -e

VERSION=${1:-dev}
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Create dist directory
mkdir -p dist

# Build for multiple platforms
PLATFORMS=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
    "windows/amd64"
)

BINARY_NAME="rc"

for PLATFORM in "${PLATFORMS[@]}"; do
    PLATFORM_SPLIT=(${PLATFORM//\// })
    GOOS=${PLATFORM_SPLIT[0]}
    GOARCH=${PLATFORM_SPLIT[1]}
    
    OUTPUT_NAME="${BINARY_NAME}"
    if [ "$GOOS" = "windows" ]; then
        OUTPUT_NAME="${BINARY_NAME}.exe"
    fi
    
    OUTPUT_PATH="dist/${BINARY_NAME}-${GOOS}-${GOARCH}"
    if [ "$GOOS" = "windows" ]; then
        OUTPUT_PATH="${OUTPUT_PATH}.exe"
    fi
    
    echo "Building for ${GOOS}/${GOARCH}..."
    
    GOOS=${GOOS} GOARCH=${GOARCH} go build \
        -ldflags "-X react-cli/internal/version.Version=${VERSION} -X react-cli/internal/version.BuildTime=${BUILD_TIME} -X react-cli/internal/version.GitCommit=${GIT_COMMIT}" \
        -o "${OUTPUT_PATH}" \
        .
    
    echo "Built: ${OUTPUT_PATH}"
done

echo ""
echo "Build complete! Binaries are in the dist/ directory"
echo "Version: ${VERSION}"
echo "Build time: ${BUILD_TIME}"
echo "Git commit: ${GIT_COMMIT}"

