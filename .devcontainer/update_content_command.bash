#!/bin/bash

set -eux -o pipefail

git -C ~/.rbenv pull
git -C ~/.rbenv/plugins/ruby-build pull

sudo apt-get update -qq
sudo apt-get dist-upgrade -y

# https://github.com/Schniz/fnm/blob/v1.38.1/README.md?plain=1#L35-L39
curl --fail --show-error --location https://fnm.vercel.app/install | bash -s -- --skip-shell
bash -ic "fnm install"

bash -ic "npm install --global prettier"
