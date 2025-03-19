#!/bin/bash

set -eux -o pipefail

git -C ~/.rbenv pull
git -C ~/.rbenv/plugins/ruby-build pull

sudo apt-get update -qq
sudo apt-get dist-upgrade -y
