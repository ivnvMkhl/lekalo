
#!/bin/bash
set -e

rm -rf ./npm_package/bin
mkdir -p ./npm_package/bin/{linux-x64,linux-arm64,darwin-arm64,darwin-x64,win32-x64}

# Сборка для Linux
GOOS=linux GOARCH=amd64 go build -o ./npm_package/bin/linux-x64/lekalo
GOOS=linux GOARCH=arm64 go build -o ./npm_package/bin/linux-arm64/lekalo
chmod +x ./npm_package/bin/linux-*/lekalo

# Сборка для macOS
GOOS=darwin GOARCH=amd64 go build -o ./npm_package/bin/darwin-x64/lekalo
GOOS=darwin GOARCH=arm64 go build -o ./npm_package/bin/darwin-arm64/lekalo
chmod +x ./npm_package/bin/darwin-*/lekalo

# Сборка для Windows
GOOS=windows GOARCH=amd64 go build -o ./npm_package/bin/win32-x64/lekalo.exe

echo "Done"
