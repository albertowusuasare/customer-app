#!/bin/sh
echo "Fetching golangci-lint from github ..."
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
echo "Running golangci-lint"
cmd="golangci-lint run"
$cmd
status=$?
## take some decision ## 
[ $status -eq 0 ] && echo "Lint run was successful :)" || echo "Lint run failed :("