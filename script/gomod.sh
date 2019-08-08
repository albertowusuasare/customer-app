#!/bin/sh
export GO111MODULE=on
go mod init $MOD_PKG_NAME
go mod tidy 