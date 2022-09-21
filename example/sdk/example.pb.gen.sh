#!/bin/bash -e

go generate ./...

# we should lint the proto
buf lint

# we should generate the package
buf generate
buf generate --template buf.gen.tag.yaml

# format
go fmt ./...
