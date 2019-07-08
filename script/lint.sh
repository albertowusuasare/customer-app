#!/bin/sh
echo "Fetching golangci-lint from github ..."
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
echo "Running golangci-lint"
golangci-lint run