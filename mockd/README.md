## 使用说明

### 1.构建与下载
目录下有release_linux.sh和release_mac.sh两个脚本，可以用来构建mockd.tar.gz压缩包，内含示例配置和对应平台的二进制文件。

**注意构建需要go环境，且安装了upx软件。**

可以在 [release](https://github.com/hankmew/tools/releases) 下载

## 2.配置说明
在指定配置目录下，创建`xxx.json`文件（必须以.json结尾）。文件内容可以复制demo.json之后再修改。

```json
{
  "quartie": [ //响应分位，顺序必须从小到大
    {"percent":10, "sleep":50}, //10%的请求会在50ms内返回
    {"percent":50, "sleep":200}, 
    {"percent":75, "sleep":300},
    {"percent":90, "sleep":400},
    {"percent":99, "sleep":1200} //99%的请求会在1200ms内返回
  ],
  "path": "/demo",  //path
  "res": {          //访问此path后返回的内容，这个{}内的内容可以自定义
    "errno": 0,
    "errmsg": "success",
    "data": {
      "a": "b"
    }
  }
}
```


## 3.启动说明
首先确保二进制文件mockd有执行权限，没有则执行 `chmod +x mockd`

```sh
./mockd [-p 8179] -c [./config]
    -p 自定义监听的端口，默认8179
    -c 自定义配置目录，默认./config/
```

使用默认端口和默认配置路径
```sh
./mockd
```

使用指定端口和配置路径
```sh
./mockd -p 8080 -c /etc/mockd
```

## 4.请求示例
访问
```sh
curl http://127.0.0.1:8179/demo
```

返回
```json
{
  "errno": 0,
    "errmsg": "success",
    "data": {
      "a": "b"
  }
}
```