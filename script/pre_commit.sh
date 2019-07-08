#!/bin/sh
echo "---- Running lint.sh ----"
echo
./script/lint.sh
echo
echo
echo "---- Running unit and integration tests ----"
go test -v github.com/albertowusuasare/customer-app/... 