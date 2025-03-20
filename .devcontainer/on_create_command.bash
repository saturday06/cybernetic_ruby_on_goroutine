#!/bin/bash

set -eux -o pipefail

cd "$(dirname "$0")"

./update_content_command.bash

latest_cruby_version=$(bash -ic "rbenv install --list-all" | grep -E '^[0-9]+\.[0-9]+\.[0-9]+$' | tail -n 1)
bash -ic "rbenv install '$latest_cruby_version'"
bash -ic "rbenv local '$latest_cruby_version'"
bash -ic "bundle"
