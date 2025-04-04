#!/bin/bash

# ref: https://freshman.tech/linting-golang/

# In a bash script, set -e is a command that enables the "exit immediately" option. When this option is set, the script will terminate immediately if any command within the script exits with a non-zero status (indicating an error).
set -e

# https://github.com/mgechev/revive
revive -config revive-config.toml -formatter friendly ./...

# https://github.com/dominikh/go-tools
staticcheck ./...

# https://golangci-lint.run/usage/linters/
# https://golangci-lint.run/usage/configuration/
# https://golangci-lint.run/usage/quick-start/
golangci-lint run ./...

# https://github.com/kisielk/errcheck
errcheck ./...
