#!/bin/bash
rm -rf ./build_bin
mkdir ./build_bin
GOOS=linux GOARCH=amd64 go build -o build_bin/lekalo_linux_amd64
GOOS=linux GOARCH=arm64 go build -o build_bin/lekalo_linux_arm64
GOOS=darwin GOARCH=arm64 go build -o build_bin/lekalo_mac_arm64
GOOS=windows GOARCH=amd64 go build -o build_bin/lekalo_win_amd64

chmod 755 ./build_bin/lekalo_linux_amd64
chmod 755 ./build_bin/lekalo_linux_arm64
chmod 755 ./build_bin/lekalo_mac_arm64
chmod 755 ./build_bin/lekalo_win_amd64

echo "Done"
