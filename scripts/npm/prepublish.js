#!/usr/bin/env node

// Pre-publish script
// Validates that binaries are built before publishing

const fs = require('fs');
const path = require('path');

const binDir = path.join(__dirname, '../../bin');

function checkBinaries() {
  if (!fs.existsSync(binDir)) {
    console.error('Error: bin/ directory does not exist');
    console.error('Please build binaries before publishing');
    process.exit(1);
  }
  
  const requiredBinaries = [
    'rc-darwin-amd64',
    'rc-darwin-arm64',
    'rc-linux-amd64',
    'rc-linux-arm64',
    'rc-windows-amd64.exe'
  ];
  
  const missing = [];
  for (const binary of requiredBinaries) {
    const binaryPath = path.join(binDir, binary);
    if (!fs.existsSync(binaryPath)) {
      missing.push(binary);
    }
  }
  
  if (missing.length > 0) {
    console.error('Error: Missing required binaries:');
    missing.forEach(b => console.error(`  - ${b}`));
    console.error('\nPlease build all binaries before publishing');
    console.error('Run: ./scripts/build.sh');
    process.exit(1);
  }
  
  console.log('✓ All binaries present');
}

checkBinaries();

