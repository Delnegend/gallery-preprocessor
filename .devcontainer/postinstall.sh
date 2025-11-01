#!/usr/bin/env bash

# just
ver="1.43.0"
curl -L -o /tmp/just.tar.gz "https://github.com/casey/just/releases/download/$ver/just-$ver-x86_64-unknown-linux-musl.tar.gz"
checksum="a1bc93654f31669fd964ea3011a5e5e9676b9b6f8adcd762606e5140632ea72d"
if [ ! "$(sha256sum /tmp/just.tar.gz | awk '{print $1}')" = "$checksum" ]; then
    echo "just checksum failed"
else
    sudo tar -xf /tmp/just.tar.gz -C /usr/local/bin just
fi
rm -f /tmp/just.tar.gz
echo 'alias j=just' >> ~/.bashrc
echo 'eval "$(just --completions bash)"' >> ~/.bashrc
echo 'complete -F _just j' >> ~/.bashrc

# go
ver="1.25.3"
curl -L -o /tmp/go.tar.gz "https://go.dev/dl/go$ver.linux-amd64.tar.gz"
checksum="0335f314b6e7bfe08c3d0cfaa7c19db961b7b99fb20be62b0a826c992ad14e0f"
if [ ! "$(sha256sum /tmp/go.tar.gz | awk '{print $1}')" = "$checksum" ]; then
    echo "go checksum failed"
else
    sudo tar -xf /tmp/go.tar.gz -C /usr/local go
fi
rm -f /tmp/go.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export PATH=$PATH:~/go/bin' >> ~/.bashrc
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:~/go/bin

go install github.com/wailsapp/wails/v2/cmd/wails@latest

# bun
curl -fsSL https://bun.sh/install | bash
echo 'export BUN_INSTALL="$HOME/.bun"' >> ~/.bashrc
echo 'export PATH="$BUN_INSTALL/bin:$PATH"' >> ~/.bashrc