#!/bin/bash

set -eux -o pipefail

cd "$(dirname "$0")"

./update_content_command.bash

bash -ic "rbenv install"
