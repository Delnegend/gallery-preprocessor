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

# go
if [ -f /usr/local/go/bin/go ]; then
    echo "Go is already installed."
else
    echo "Installing Go..."
    wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz -O go.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go.tar.gz
    rm go.tar.gz
fi

# add /usr/local/go/bin to path
if [[ ":$PATH:" != *":/usr/local/go/bin:"* ]]; then
    echo "Adding /usr/local/go/bin to PATH..."
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
    source ~/.zshrc
fi

# add ~/go/bin to path
if [[ ":$PATH:" != *":~/go/bin:"* ]]; then
    echo "Adding ~/go/bin to PATH..."
    echo 'export PATH=$PATH:~/go/bin' >> ~/.zshrc
    source ~/.zshrc
fi

# wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest