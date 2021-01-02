# 交叉编译为linux x86平台的可执行程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w"
# 使用upx压缩，注意需要先安装upx
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 upx -9 mockd