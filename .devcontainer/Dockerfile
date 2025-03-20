# SPDX-License-Identifier: MIT
FROM golang:1-bookworm

# https://github.com/hadolint/hadolint/wiki/DL4006
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# https://github.com/rbenv/ruby-build/wiki#ubuntudebianmint
RUN apt-get update -q \
    && apt-get install \
        'git=*' \
        'sudo=*' \
        'autoconf=*' \
        'patch=*' \
        'build-essential=*' \
        'rustc=*' \
        'libssl-dev=*' \
        'libyaml-dev=*' \
        'libreadline6-dev=*' \
        'shellcheck=*' \
        'shfmt=*' \
        'zlib1g-dev=*' \
        'libgmp-dev=*' \
        'libncurses5-dev=*' \
        'libffi-dev=*' \
        'libgdbm-dev=*' \
        'libdb-dev=*' \
        'uuid-dev=*' \
        --yes --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN useradd --create-home --user-group --shell /bin/bash developer \
    && usermod --append --groups sudo developer
RUN mkdir -p /workspace && chown developer:developer /workspace
USER developer
WORKDIR /home/developer

# https://github.com/rbenv/rbenv/blob/v1.3.2/README.md#basic-git-checkout
RUN git clone https://github.com/rbenv/rbenv.git ~/.rbenv && ~/.rbenv/bin/rbenv init

# https://github.com/rbenv/ruby-build/tree/v20250318?tab=readme-ov-file#clone-as-rbenv-plugin-using-git
RUN git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
