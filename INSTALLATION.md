# Installation Setup Guide

This document explains how to set up the installer system for your React CLI tool.

## Before First Release

### 1. Update Repository URLs

Update the following files with your actual GitHub username and repository name:

**Files updated:**
- ✅ `install.sh` - Updated to `khotcholava/zhvabu-cli`
- ✅ `install.ps1` - Updated to `khotcholava/zhvabu-cli`
- ✅ `package.json` - Updated repository URL
- ✅ `Formula/react-cli.rb` - Updated download URLs
- ✅ `.github/workflows/release.yml` - Uses `${{ github.repository }}` automatically
- ✅ `README.md` - Updated all installation URLs

### 2. Create GitHub Repository

1. Create a new repository on GitHub
2. Push your code
3. Update all URLs in the files above

### 3. First Release

1. **Tag your first release:**
   ```bash
   git tag v0.1.0
   git push origin v0.1.0
   ```

2. **GitHub Actions will automatically:**
   - Build binaries for all platforms
   - Create a GitHub Release
   - Upload binaries as release assets

3. **Update Homebrew formula SHA256:**
   After the first release, get the SHA256 hashes:
   ```bash
   shasum -a 256 dist/rc-darwin-amd64.tar.gz
   shasum -a 256 dist/rc-darwin-arm64.tar.gz
   shasum -a 256 dist/rc-linux-amd64.tar.gz
   shasum -a 256 dist/rc-linux-arm64.tar.gz
   ```
   Update `Formula/react-cli.rb` with these values.

### 4. Homebrew Tap Setup (Optional)

To enable `brew install`:

1. Create a new repository: `homebrew-tap`
2. Copy `Formula/react-cli.rb` to that repository
3. Update the formula with correct URLs and SHA256
4. Users can then:
   ```bash
   brew tap khotcholava/homebrew-tap
   brew install react-cli
   ```

### 5. npm Publishing (Optional)

To publish to npm:

1. Update `package.json` with your details
2. Create npm account if needed
3. Build binaries and place in `bin/` directory
4. Publish:
   ```bash
   npm publish
   ```

## Testing Installation

### Test Install Script Locally

1. Build binaries:
   ```bash
   ./scripts/build.sh v0.1.0
   ```

2. Create a test release (or use existing):
   - Upload binaries manually to a GitHub Release
   - Or test with local files

3. Test install script:
   ```bash
   # Update install.sh to point to test release
   bash install.sh v0.1.0
   ```

## Release Process

1. **Update version:**
   - Update version in `internal/version/version.go` (or use build flags)
   - Update version in `package.json`
   - Update version in `Formula/react-cli.rb`

2. **Commit and tag:**
   ```bash
   git add .
   git commit -m "chore: release v0.1.0"
   git tag v0.1.0
   git push origin main
   git push origin v0.1.0
   ```

3. **GitHub Actions will:**
   - Build all binaries
   - Create release
   - Upload assets

4. **Update Homebrew formula** (if using):
   - Get SHA256 hashes from release
   - Update formula
   - Commit to tap repository

## Verification

After release, verify installation works:

```bash
# Test install script
curl -fsSL https://raw.githubusercontent.com/khotcholava/zhvabu-cli/main/install.sh | sh

# Verify installation
rc --version
rc version
```

## Troubleshooting

### Build fails
- Check Go version (requires 1.24+)
- Verify all dependencies are installed

### Install script fails
- Check GitHub repository URL is correct
- Verify release exists and binaries are uploaded
- Check network connectivity

### Homebrew install fails
- Verify SHA256 hashes are correct
- Check formula URLs point to correct release
- Ensure tap repository is set up correctly

