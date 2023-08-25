#!/bin/bash

# ref: https://blog.devgenius.io/sort-go-imports-acb76224dfa7
# https://yolken.net/blog/cleaner-go-code-golines

# In a bash script, set -e is a command that enables the "exit immediately" option. When this option is set, the script will terminate immediately if any command within the script exits with a non-zero status (indicating an error).
set -e

# https://github.com/mvdan/gofumpt
gofumpt -l -w .

# https://golang.org/cmd/gofmt/
# gofmt -w .

# # https://pkg.go.dev/golang.org/x/tools/cmd/goimports
# goimports -l -w .

# https://github.com/incu6us/goimports-reviser
# will do `gofmt` and `goimports` internally
# -rm-unused, -set-alias have some errors
goimports-reviser -local -format -recursive ./...

# https://github.com/segmentio/golines
golines -m 120 -w --ignore-generated .

