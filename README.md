Golang 服务器返回公网地址

```
git地址: https://github.com/LoongLiWang/DockerFIle_ReturnOutIP
```

> Fun

## 客户端相关
### 扩展

#### shell
```shell
# curl ip.wang-li.top:93
```

#### python客户端
``` python
#!/usr/bin/env python3

import requests

def main():
    url = "http://ip.wang-li.top:93/4u6385IP"
    MyIP = requests.get(url).text
    print(MyIP)

if __name__ == '__main__':
    main()
```

#### php
```php
<?php
	$url='http://ip.wang-li.top:93';

	$result=file_get_contents($url);

	echo $result
?>
```

#### Golang客户端
``` golang
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "URL"

	resp , err := http.Get(url);if err != nil {
		fmt.Println("Http Connect Error", err)
	} else {
		b , err := ioutil.ReadAll(resp.Body); if err != nil {
			fmt.Println("Read Body Error" , err)
		} else {
			fmt.Printf("%s",b)
		}
	}
}
```

#### Lua
```lua
#!/usr/bin/lua5.3

http = require("socket.http")

function ReturnIP(url)
	local resp = http.request(url)

	return resp
end

print(ReturnIP('http://ip.wang-li.top:93/4u6385IP'))
```

## 服务器相关
### Docker部署:
Docker Hub 地址:  https://hub.docker.com/r/2859413527/retuen_out_ip
# 返回公网IP地址 Docker 

## 下载镜像

```sh
# docker pull 2859413527/retuen_out_ip
```



## 最简单运行容器

```sh
# docker run -d --name rcs -p 93:93 2859413527/retuen_out_ip
```



## 定义更多的环境变量

### 定义路由

默认值: 4u6385IP

```sh
# docker run -d --name rcs -p 93:93 -e ListenRoute="/dev" 2859413527/retuen_out_ip
```

含义: 定义新的路由 /dev 



### 定义限流时长

默认值: 60(单位: 秒)

```sh
# docker run -d --name rcs -p 93:93 -e LimitTime=60 2859413527/retuen_out_ip
```

含义: 定义统计限流时长 60秒



### 定义限流个数

默认值: 60(单位: 次)

```sh
# docker run -d --name rcs -p 93:93 -e LimitCount=10 2859413527/retuen_out_ip
```

含义: 定义在限流时长内，每个公网客户端限流为10次



## 例子

定义 路由 /dev1 , 限流时长为 120秒，限流个数为 60 次
```sh
# docker run -d --name rcs -e ListenRoute="/dev1" -e LimitTime=120 -e LimitCount=60 -p 93:93 2859413527/retuen_out_ip
```

### 获取帮助:
``` bash
# ./ReturnOutIP -h
Usage of ./ReturnOutIP:
  -LimitCount int
        Set the  flow limit throughput within the time (default 10)
  -LimitTime int
        Set flow limit time , (Second) (default 1)
  -ListenAddr string
        Set http server listen address (default "0.0.0.0:93")
  -ListenRoute string
        Set http server listen Route (default "/4u6385IP")
  -MongoAuthDB string
        Mongo Auth DB
  -MongoHost string
        Mongo Host (default "127.0.0.1")
  -MongoOn int
        Mongo Flag (0: Off 1:On)
  -MongoPass string
        Mongo Password
  -MongoUser string
        MongoUser
  -h    This help
#
```

### 默认参数
默认运行监听的端口为: 0.0.0.0:93 监听的http路由为 /4u6385IP
``` bash
# ./ReturnOutIP
2020/07/10 11:49:23 LimitTime: 60 's   LimitCount: 50
2020/07/10 11:49:23 Server running on http://0.0.0.0:93
2020/07/10 11:49:23 Server running on http://0.0.0.0:93/4u6385IP

```

### 指定监听的端口和路由
``` bash
# ./ReturnOutIP -ListenAddr "0.0.0.0:95" -ListenRoute '/OutIP'
2020/07/10 13:27:20 LimitTime: 1 's   LimitCount: 10000
2020/07/10 13:27:20 Server running on http://0.0.0.0:95
2020/07/10 13:27:20 Server running on http://0.0.0.0:95/OutIP
```

### 限流
-LimitTime: 间隔时间 -LimitCount:在间隔时间内允许流入的量
比如: 60s内只允许10个流入
```bash
./ReturnOutIP -LimitTime 60 -LimitCount 10
```

### 日志Mongo落地
```bash
# ./ReturnOutIP -MongoHost=172.27.0.12 -MongoAuthDB=admin -MongoOn=1 -MongoUser=mongouser -MongoPass=123456
```