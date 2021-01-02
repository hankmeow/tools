mkdir -p release/mockd
cp -r config release/mockd/

# 交叉编译为linux x86平台的可执行程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -o release/mockd/mockd
# 使用upx压缩，注意需要先安装upx
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 upx -9 release/mockd/mockd

cd release
chmod +x mockd/mockd
tar zcvf mockd-linux.tar.gz mockd
rm -rf mockd