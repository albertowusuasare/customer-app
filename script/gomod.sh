#!/bin/sh

function CreateModule(){
    set -e
    export GO111MODULE=on
    PKG_NAME=$1 # PKG_NAME is the first CLI arg
    go mod init $PKG_NAME
    go mod tidy 
}