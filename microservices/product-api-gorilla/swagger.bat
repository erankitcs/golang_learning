@echo off
echo.
docker run --rm -it --env GOPATH=/go -v %CD%:/go/src -w /go/src quay.io/goswagger/swagger generate spec -o ./swagger.yaml --scan-models