#!/usr/bin/env bash

# zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended

# node
curl https://get.volta.sh | bash
export VOLTA_HOME="$HOME/.volta" && export PATH="$VOLTA_HOME/bin:$PATH"
echo 'export VOLTA_HOME="$HOME/.volta"' >> ~/.zshrc
echo 'export PATH="$VOLTA_HOME/bin:$PATH"' >> ~/.zshrc
volta install node@lts pnpm && pnpm config set store-dir ~/.pnpm-store
cd frontend && pnpm install && cd ..

# just
ver="1.42.3"
curl -L -o /tmp/just.tar.gz https://github.com/casey/just/releases/download/$ver/just-$ver-x86_64-unknown-linux-musl.tar.gz
checksum=$(openssl dgst -sha3-512 /tmp/just.tar.gz | awk '{print $2}')
expected="939eef7b9105e4805825b6cf7aac7bb2e8daf1aa07e4e8cfce619e1e447a444847318abd0c7e0ac3553e80a98e8d465158c7574e1fc7d2232e839b021f9dad67"
if [ ! "$checksum" = "$expected" ]; then
    echo "just tarball checksum failed\nexpected: $expected\ngot: $checksum"
else
    sudo rm -rf /usr/local/bin/just
    sudo tar -xf /tmp/just.tar.gz -C /usr/local/bin just
fi
rm -f /tmp/just.tar.gz
echo 'alias j=just' >> ~/.zshrc
just --completions zsh > ~/.just.zsh
echo '[[ -f ~/.just.zsh ]] && source ~/.just.zsh' >> ~/.zshrc

# go
if [ -f /usr/local/go/bin/go ]; then
    echo "Go is already installed."
else
    echo "Installing Go..."
    wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz -O go.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go.tar.gz
    rm go.tar.gz
fi

# PATH stuffs
if [[ ":$PATH:" != *":/usr/local/go/bin:"* ]]; then
    echo "Adding /usr/local/go/bin to PATH..."
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
    source ~/.zshrc
fi
if [[ ":$PATH:" != *":~/go/bin:"* ]]; then
    echo "Adding ~/go/bin to PATH..."
    echo 'export PATH=$PATH:~/go/bin' >> ~/.zshrc
    source ~/.zshrc
fi

go install github.com/wailsapp/wails/v2/cmd/wails@latest