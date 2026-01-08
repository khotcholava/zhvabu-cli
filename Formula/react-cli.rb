# Homebrew formula for react-cli
# To use this formula, create a tap:
#   brew tap khotcholava/homebrew-tap
#   brew install react-cli

class ReactCli < Formula
  desc "React CLI - Generate React components with ease"
  homepage "https://github.com/khotcholava/zhvabu-cli"
  version "0.3.0"
  license "MIT"

  # Update these URLs with your actual GitHub repository
  if OS.mac? && Hardware::CPU.intel?
    url "https://github.com/khotcholava/zhvabu-cli/releases/download/v#{version}/rc-darwin-amd64.tar.gz"
    sha256 "" # Add SHA256 after first release
  elsif OS.mac? && Hardware::CPU.arm?
    url "https://github.com/khotcholava/zhvabu-cli/releases/download/v#{version}/rc-darwin-arm64.tar.gz"
    sha256 "" # Add SHA256 after first release
  elsif OS.linux? && Hardware::CPU.intel?
    url "https://github.com/khotcholava/zhvabu-cli/releases/download/v#{version}/rc-linux-amd64.tar.gz"
    sha256 "" # Add SHA256 after first release
  elsif OS.linux? && Hardware::CPU.arm?
    url "https://github.com/khotcholava/zhvabu-cli/releases/download/v#{version}/rc-linux-arm64.tar.gz"
    sha256 "" # Add SHA256 after first release
  end

  def install
    bin.install "rc"
  end

  test do
    system "#{bin}/rc", "--version"
  end
end

