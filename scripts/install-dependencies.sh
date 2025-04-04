#!/bin/bash

# In a bash script, set -e is a command that enables the "exit immediately" option. When this option is set, the script will terminate immediately if any command within the script exits with a non-zero status (indicating an error).
set -e

# https://www.reddit.com/r/golang/comments/x722i0/go_install_vs_go_mod_tidy_vs_go_get/
# will do `go mod download` implicitly - Adds missing dependencies - Removes unused dependencies - Update go.mod and go.sum files
go mod tidy



