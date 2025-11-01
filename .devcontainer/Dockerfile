FROM debian:12-slim

RUN apt-get update && apt-get install -y \
    build-essential libgtk-3-dev libwebkit2gtk-4.0-dev pkg-config \
    curl git sudo unzip && apt-get clean

RUN groupadd -g 1000 gallery-pp && \
    useradd -m -u 1000 -g gallery-pp gallery-pp && \
    echo "gallery-pp ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/gallery-pp

USER gallery-pp
