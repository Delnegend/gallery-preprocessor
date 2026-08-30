FROM debian:13-slim

ARG GO_VERSION=1.27.0
ARG GO_ARCH=linux-amd64

ENV LANG=C.UTF-8
ENV LC_ALL=C.UTF-8
ENV SHELL=/bin/bash

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y --no-install-recommends \
    build-essential ca-certificates curl git just fzf \
    libgtk-3-dev libwebkit2gtk-4.1-dev pkg-config sudo unzip && \
    rm -rf /var/lib/apt/lists/*

# Go toolchain from the official tarball (pinned to go.mod's version).
RUN curl -fsSL -o /tmp/go.tgz \
        "https://go.dev/dl/go${GO_VERSION}.${GO_ARCH}.tar.gz" && \
    tar -C /usr/local -xzf /tmp/go.tgz && \
    rm /tmp/go.tgz

ENV PATH="/home/gallery/go/bin:/usr/local/go/bin:${PATH}"

RUN groupadd -g 1000 gallery && \
    useradd -m -u 1000 -g gallery -s /bin/bash gallery && \
    echo "gallery ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/gallery && \
    chmod 0440 /etc/sudoers.d/gallery && \
    chsh -s /bin/bash gallery

SHELL ["/bin/bash", "-c"]
USER gallery

# Pre-install wails and bun so container works even if postCreateCommand is skipped (e.g. Zed)
# wails version is pinned to go.mod (v2.15.0) to avoid mismatch warning and respect 2-week min-age
RUN go install github.com/wailsapp/wails/v2/cmd/wails@v2.15.0
RUN curl -fsSL https://bun.sh/install | bash && sudo ln -sf /home/gallery/.bun/bin/bun /usr/local/bin/bun

# just completions for interactive shells (was in postinstall.sh, now in image for cache)
RUN echo 'eval "$(just --completions bash)"' >> /home/gallery/.bashrc

# Ensure wails/bun are in PATH for non-interactive shells (Zed, `just build`)
ENV PATH="/home/gallery/.bun/bin:/home/gallery/go/bin:/usr/local/go/bin:${PATH}"
