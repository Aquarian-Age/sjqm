#!/bin/bash

# set -v
# go get -u github.com/ying32/liblclbinres

go build -o sjqm-linux-x86_64 -i -tags tempdll -ldflags="-s -w" .
GOOS=windows GOARCH=amd64 go build -o sjqm-windows-x86_64.exe -i -tags tempdll -ldflags="-H windowsgui -s -w" .
GOOS=windows GOARCH=386 go build -o sjqm-windows-x86.exe -i -tags tempdll -ldflags="-H windowsgui -s -w" .
# go get gioui.org/cmd/gogio
# go run gioui.org/cmd/gogio -target android -ldflags="-w -s" .