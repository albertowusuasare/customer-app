#!/bin/sh
set -e
echo "Running golangci-lint"
cmd="golangci-lint run"
$cmd
status=$?

if [ $status -eq 0 ]; then
        echo "Lint run was successful :)"  
    else
        echo "Lint run failed :(" 
        exit 1
fi