# React CLI Installer for Windows
# PowerShell script to install react-cli

$ErrorActionPreference = "Stop"

# Configuration
$Repo = "khotcholava/zhvabu-cli"
$BinaryName = "rc"
$Version = if ($args.Count -gt 0) { $args[0] } else { "latest" }

# Colors
function Write-ColorOutput($ForegroundColor) {
    $fc = $host.UI.RawUI.ForegroundColor
    $host.UI.RawUI.ForegroundColor = $ForegroundColor
    if ($args) {
        Write-Output $args
    }
    $host.UI.RawUI.ForegroundColor = $fc
}

# Detect platform
function Get-Platform {
    $arch = (Get-WmiObject Win32_Processor).AddressWidth
    
    if ($arch -eq 64) {
        return "windows/amd64"
    } else {
        Write-ColorOutput Red "32-bit Windows is not supported"
        exit 1
    }
}

# Get latest version from GitHub API
function Get-LatestVersion {
    if ($Version -eq "latest") {
        $response = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
        return $response.tag_name
    } else {
        return $Version
    }
}

# Download and install binary
function Install-Binary {
    $platform = Get-Platform
    $os = $platform.Split('/')[0]
    $arch = $platform.Split('/')[1]
    
    $versionTag = Get-LatestVersion
    
    Write-ColorOutput Green "Installing $BinaryName $versionTag for $os/$arch..."
    
    $downloadUrl = "https://github.com/$Repo/releases/download/$versionTag/${BinaryName}-${os}-${arch}.zip"
    
    $tempDir = New-TemporaryFile | ForEach-Object { Remove-Item $_; New-Item -ItemType Directory -Path $_ }
    
    try {
        Write-ColorOutput Yellow "Downloading from $downloadUrl..."
        
        $zipPath = Join-Path $tempDir "$BinaryName.zip"
        Invoke-WebRequest -Uri $downloadUrl -OutFile $zipPath -UseBasicParsing
        
        Write-ColorOutput Yellow "Extracting..."
        Expand-Archive -Path $zipPath -DestinationPath $tempDir -Force
        
        $binaryPath = Join-Path $tempDir "$BinaryName.exe"
        
        if (-not (Test-Path $binaryPath)) {
            Write-ColorOutput Red "Binary not found in archive"
            exit 1
        }
        
        # Determine install location
        $localBin = Join-Path $env:USERPROFILE "bin"
        if (-not (Test-Path $localBin)) {
            New-Item -ItemType Directory -Path $localBin -Force | Out-Null
        }
        
        $installPath = Join-Path $localBin "$BinaryName.exe"
        
        Write-ColorOutput Yellow "Installing to $installPath..."
        Copy-Item -Path $binaryPath -Destination $installPath -Force
        
        # Add to PATH if not already there
        $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
        if ($currentPath -notlike "*$localBin*") {
            Write-ColorOutput Yellow "Adding $localBin to PATH..."
            [Environment]::SetEnvironmentVariable("Path", "$currentPath;$localBin", "User")
            $env:Path += ";$localBin"
        }
        
        # Verify installation
        if (Get-Command $BinaryName -ErrorAction SilentlyContinue) {
            $installedVersion = & $BinaryName --version 2>&1
            Write-ColorOutput Green "Successfully installed $BinaryName $installedVersion"
            Write-ColorOutput Green "Run '$BinaryName --help' to get started"
            Write-ColorOutput Yellow "Note: You may need to restart your terminal for PATH changes to take effect"
        } else {
            Write-ColorOutput Yellow "Installation completed, but $BinaryName is not in PATH"
            Write-ColorOutput Yellow "Add $localBin to your PATH or restart your terminal"
        }
    } finally {
        Remove-Item -Recurse -Force $tempDir -ErrorAction SilentlyContinue
    }
}

# Main
function Main {
    Write-ColorOutput Green "React CLI Installer"
    Write-Output ""
    
    # Check for required commands
    if (-not (Get-Command Invoke-WebRequest -ErrorAction SilentlyContinue)) {
        Write-ColorOutput Red "PowerShell 3.0 or higher is required"
        exit 1
    }
    
    Install-Binary
}

Main

