Golang 服务器返回公网地址

```
gitee地址: https://gitee.com/wang_li/ReturnOutIP
```

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
### 获取帮助:
``` bash
# ./ReturnOutIP -h
Usage of ./ReturnOutIP:
  -ListenAddr string
        Set http server listen address (default "0.0.0.0:93")
  -ListenRoute string
        Set http server listen Route (default "/4u6385IP")
  -h    This help
#
```

### 默认参数
默认运行监听的端口为: 0.0.0.0:93 监听的http路由为 /4u6385IP
``` bash
# ./ReturnOutIP
2019/09/03 15:36:55 Server running on http://0.0.0.0:93/4u6385IP
```

### 指定监听的端口和路由
``` bash
# ./ReturnOutIP -ListenAddr "0.0.0.0:95" -ListenRoute '/OutIP'
2019/09/03 15:41:24 Server running on http://0.0.0.0:95/OutIP
```

### 限流
-LimitTime: 间隔时间 -LimitCount:在间隔时间内允许流入的量
比如: 60s内只允许10个流入
```bash
./ReturnOutIP -LimitTime 60 -LimitCount 10
```

#### 例子
```zsh
~/.../src/ReturnOutIP >>> sudo ./ReturnOutIP -LimitTime 60 -LimitCount 10                                                                           ±[●●][master]
[sudo] liwang 的密码：
2020/07/09 21:35:10 LimitTime: 60 's   LimitCount: 10
2020/07/09 21:35:10 Server running on http://0.0.0.0:93
2020/07/09 21:35:10 Server running on http://0.0.0.0:93/4u6385IP
2020/07/09 21:35:17 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 0 0 0}
2020/07/09 21:35:17 2020-07-09 21:35:17.573268239 +0800 CST m=+6.926567676 1 : 正常
2020/07/09 21:35:17 RemoteAddr: 127.0.0.1 URL: /favicon.ico UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301717 1}
2020/07/09 21:35:17 2020-07-09 21:35:17.644367798 +0800 CST m=+6.997667216 2 : 正常
2020/07/09 21:35:46 RemoteAddr: 192.168.3.8 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{192.168.3.8 0 0 0}
2020/07/09 21:35:46 2020-07-09 21:35:46.527050739 +0800 CST m=+35.880350172 1 : 正常
2020/07/09 21:35:46 RemoteAddr: 192.168.3.8 URL: /favicon.ico UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{192.168.3.8 1594301746 1594301746 1}
2020/07/09 21:35:46 2020-07-09 21:35:46.702135218 +0800 CST m=+36.055434646 2 : 正常
2020/07/09 21:35:48 RemoteAddr: 192.168.3.8 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{192.168.3.8 1594301746 1594301746 2}
2020/07/09 21:35:48 2020-07-09 21:35:48.582040932 +0800 CST m=+37.935340367 3 : 正常
2020/07/09 21:35:49 RemoteAddr: 192.168.3.8 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{192.168.3.8 1594301746 1594301748 3}
2020/07/09 21:35:49 2020-07-09 21:35:49.91972562 +0800 CST m=+39.273025130 4 : 正常
2020/07/09 21:35:52 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301717 2}
2020/07/09 21:35:52 2020-07-09 21:35:52.69996333 +0800 CST m=+42.053262829 3 : 正常
2020/07/09 21:35:53 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301752 3}
2020/07/09 21:35:53 2020-07-09 21:35:53.98599809 +0800 CST m=+43.339297534 4 : 正常
2020/07/09 21:35:54 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301753 4}
2020/07/09 21:35:54 2020-07-09 21:35:54.700779385 +0800 CST m=+44.054078825 5 : 正常
2020/07/09 21:35:55 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301754 5}
2020/07/09 21:35:55 2020-07-09 21:35:55.332310371 +0800 CST m=+44.685609881 6 : 正常
2020/07/09 21:35:57 RemoteAddr: 192.168.3.8 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{192.168.3.8 1594301746 1594301749 4}
2020/07/09 21:35:57 2020-07-09 21:35:57.924001807 +0800 CST m=+47.277301246 5 : 正常
2020/07/09 21:35:57 RemoteAddr: 192.168.3.8 URL: /favicon.ico UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{192.168.3.8 1594301746 1594301757 5}
2020/07/09 21:35:57 2020-07-09 21:35:57.949954388 +0800 CST m=+47.303253810 6 : 正常
2020/07/09 21:36:00 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301755 6}
2020/07/09 21:36:00 2020-07-09 21:36:00.4178457 +0800 CST m=+49.771145133 7 : 正常
2020/07/09 21:36:01 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301760 7}
2020/07/09 21:36:01 2020-07-09 21:36:01.840603169 +0800 CST m=+51.193902624 8 : 正常
2020/07/09 21:36:02 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301761 8}
2020/07/09 21:36:02 2020-07-09 21:36:02.141744597 +0800 CST m=+51.495044085 9 : 正常
2020/07/09 21:36:02 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301762 9}
2020/07/09 21:36:02 2020-07-09 21:36:02.431529494 +0800 CST m=+51.784829019 10 : 正常
2020/07/09 21:36:02 RemoteAddr: 127.0.0.1 URL: / UserAgent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0
{127.0.0.1 1594301717 1594301762 10}
2020/07/09 21:36:02 2020-07-09 21:36:02.72304481 +0800 CST m=+52.076344261 11 : 限流
```

