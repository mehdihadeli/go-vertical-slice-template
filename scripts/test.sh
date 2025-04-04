#!/bin/bash

# https://blog.devgenius.io/go-golang-testing-tools-tips-to-step-up-your-game-4ed165a5b3b5
# https://github.com/testcontainers/testcontainers-go/pull/1394
# https://github.com/testcontainers/testcontainers-go/issues/1359

# In a bash script, set -e is a command that enables the "exit immediately" option. When this option is set, the script will terminate immediately if any command within the script exits with a non-zero status (indicating an error).
set -e

readonly type="$1"

 go test -tags="$type" -timeout=30m  -count=1 -p=1 -parallel=1 ./...

