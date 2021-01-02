CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 upx -9 mockd