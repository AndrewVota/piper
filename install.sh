#!/bin/bash

# Define variables
REPO="AndrewVota/piper"
TAG=$(curl --silent "https://api.github.com/repos/$REPO/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")')

# Remove the 'v' from the tag if present
# TAG=${TAG#v}

# Debug: Print the tag to ensure it's fetched correctly
echo "Latest tag fetched: $TAG"
# Exit if tag is not found
if [ -z "$TAG" ]; then
  echo "Error: Unable to fetch the latest tag."
  exit 1
fi

# Detect the platform
ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

# Map architectures to the names used in the release
case $ARCH in
  "x86_64")
    ARCH="amd64"
    ;;
  "aarch64" | "arm64")
    ARCH="arm64"
    ;;
  "i386" | "i686")
    ARCH="386"
    ;;
esac

# Special case for Darwin (macOS)
if [ "$OS" == "darwin" ]; then
  OS="Darwin"
elif [ "$OS" == "linux" ]; then
  OS="linux"
elif [[ "$OS" == "mingw"* || "$OS" == "cygwin"* ]]; then
  OS="windows"
fi

# Download the appropriate binary
URL="https://github.com/$REPO/releases/download/$TAG/piper_${TAG#v}_${OS}_${ARCH}.tar.gz"
echo "Downloading from URL: $URL"
curl -L $URL -o piper_${OS}_${ARCH}.tar.gz

# Check if the downloaded file is a valid tar.gz file
if file piper_${OS}_${ARCH}.tar.gz | grep -q 'gzip compressed data'; then
  # Extract the tarball
  tar -xzf piper_${OS}_${ARCH}.tar.gz
  
  # Find the binary
  BINARY=$(find . -name piper)
  
  # Move the binary to a directory in PATH
  sudo mv $BINARY /usr/local/bin/
  
  # Clean up
  rm -rf piper_${OS}_${ARCH}.tar.gz piper_${OS}_${ARCH}
  
  echo "piper has been installed successfully."
else
  echo "Downloaded file is not a valid gzip compressed file."
  rm -f piper_${OS}_${ARCH}.tar.gz
  exit 1
fi
