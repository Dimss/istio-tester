#!/usr/bin/env bash
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh;
ln -s /builds/dimss/my-test-proj/ /go/src/
cd /go/src/my-test-proj
dep ensure;
export CGO_ENABLED=0 GOOS=linux GOARCH=amd64;
ls -all
go test
go build -o istio-teste-new-build;
ls -all
