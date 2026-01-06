#!/usr/bin/env node

// Post-install script for npm package
// Downloads the appropriate binary for the platform

const fs = require("fs");
const path = require("path");
const https = require("https");
const os = require("os");

const REPO = "khotcholava/zhvabu-cli";
const VERSION = require("../../package.json").version;

function getPlatform() {
  const platform = os.platform();
  const arch = os.arch();

  let osName;
  switch (platform) {
    case "darwin":
      osName = "darwin";
      break;
    case "linux":
      osName = "linux";
      break;
    case "win32":
      osName = "windows";
      break;
    default:
      throw new Error(`Unsupported platform: ${platform}`);
  }

  let archName;
  switch (arch) {
    case "x64":
      archName = "amd64";
      break;
    case "arm64":
      archName = "arm64";
      break;
    default:
      throw new Error(`Unsupported architecture: ${arch}`);
  }

  return { os: osName, arch: archName };
}

function downloadBinary(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    https
      .get(url, (response) => {
        if (response.statusCode === 302 || response.statusCode === 301) {
          // Handle redirect
          https
            .get(response.headers.location, (redirectResponse) => {
              redirectResponse.pipe(file);
              file.on("finish", () => {
                file.close();
                fs.chmodSync(dest, 0o755);
                resolve();
              });
            })
            .on("error", reject);
        } else {
          response.pipe(file);
          file.on("finish", () => {
            file.close();
            fs.chmodSync(dest, 0o755);
            resolve();
          });
        }
      })
      .on("error", reject);
  });
}

async function main() {
  try {
    const { os: osName, arch: archName } = getPlatform();
    const ext = osName === "windows" ? "exe" : "";
    const binaryName = `rc${ext ? "." + ext : ""}`;
    const archiveExt = osName === "windows" ? "zip" : "tar.gz";

    const binDir = path.join(__dirname, "../../bin");
    if (!fs.existsSync(binDir)) {
      fs.mkdirSync(binDir, { recursive: true });
    }

    const binaryPath = path.join(binDir, binaryName);
    const downloadUrl = `https://github.com/${REPO}/releases/download/v${VERSION}/rc-${osName}-${archName}.${archiveExt}`;

    console.log(`Downloading ${binaryName} for ${osName}/${archName}...`);

    // For npm, we'd need to extract the archive
    // This is a simplified version - you might want to use a library like tar or yauzl
    console.log(
      "Note: Full binary download/extraction needs to be implemented"
    );
    console.log(`Binary should be downloaded from: ${downloadUrl}`);
  } catch (error) {
    console.error("Error during postinstall:", error.message);
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}
