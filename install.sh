#!/bin/bash

REPO="AndrewVota/piper"
TAG=$(curl --silent "https://api.github.com/repos/$REPO/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")')

ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

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

if [ "$OS" == "darwin" ]; then
  OS="Darwin"
elif [ "$OS" == "linux" ]; then
  OS="linux"
elif [[ "$OS" == "mingw"* || "$OS" == "cygwin"* ]]; then
  OS="windows"
fi

URL="https://github.com/$REPO/releases/download/$TAG/piper_${TAG#v}_${OS}_${ARCH}.tar.gz"
echo "Downloading from URL: $URL"
curl -L $URL -o piper_${OS}_${ARCH}.tar.gz

tar -xzf piper_${OS}_${ARCH}.tar.gz
BINARY=$(find . -name piper)
sudo mv $BINARY /usr/local/bin/
rm -rf piper_${OS}_${ARCH}.tar.gz piper_${OS}_${ARCH}

echo "piper has been installed successfully."

