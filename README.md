Golang 服务器返回公网地址

```
gitee地址: https://gitee.com/wang_li/ReturnOutIP
```

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

### 扩展
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