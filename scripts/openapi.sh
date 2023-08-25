#!/bin/bash

# In a bash script, set -e is a command that enables the "exit immediately" option. When this option is set, the script will terminate immediately if any command within the script exits with a non-zero status (indicating an error).
set -e

swag init --parseDependency --parseInternal --parseDepth 1  -g ./cmd/app/main.go  -d "./" -o "./docs"
swag init --parseDependency --parseInternal --parseDepth 1  -g ./cmd/app/main.go  -d "./" -o "./api/openapi/"
