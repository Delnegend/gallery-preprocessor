FROM debian:13-slim

ARG GO_VERSION=1.24.0
ARG GO_ARCH=linux-amd64

ENV LANG=C.UTF-8
ENV LC_ALL=C.UTF-8

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y --no-install-recommends \
    build-essential ca-certificates curl git just \
    libgtk-3-dev libwebkit2gtk-4.1-dev pkg-config sudo unzip && \
    rm -rf /var/lib/apt/lists/*

# Go toolchain from the official tarball (pinned to go.mod's version).
RUN curl -fsSL -o /tmp/go.tgz \
        "https://go.dev/dl/go${GO_VERSION}.${GO_ARCH}.tar.gz" && \
    tar -C /usr/local -xzf /tmp/go.tgz && \
    rm /tmp/go.tgz

ENV PATH="/home/gallery-pp/go/bin:/usr/local/go/bin:${PATH}"

RUN groupadd -g 1000 gallery-pp && \
    useradd -m -u 1000 -g gallery-pp gallery-pp && \
    echo "gallery-pp ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/gallery-pp && \
    chmod 0440 /etc/sudoers.d/gallery-pp

USER gallery-pp
