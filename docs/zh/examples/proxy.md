# 代理示例

## HTTP 代理

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
		Proxy: "http://127.0.0.1:8080",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```

## HTTPS 代理

```go
response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
	Proxy: "https://127.0.0.1:8080",
})
```

## SOCKS5 代理

```go
response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
	Proxy: "socks5://127.0.0.1:1080",
})
```

## 带认证的代理

```go
response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
	Proxy: "http://username:password@127.0.0.1:8080",
})
```

## 使用 Fetch 实例

```go
f := fetch.New()
f.SetProxy("http://127.0.0.1:8080")

response, err := f.Get("https://httpbin.zcorky.com/ip").Execute()
```

## 环境变量

Fetch 自动使用环境变量中的代理设置：

```bash
export HTTP_PROXY=http://127.0.0.1:8080
export HTTPS_PROXY=http://127.0.0.1:8080
export SOCKS_PROXY=socks5://127.0.0.1:1080
```

无需在代码中显式配置：

```go
// 自动使用 HTTP_PROXY/HTTPS_PROXY/SOCKS_PROXY
response, err := fetch.Get("https://httpbin.zcorky.com/ip")
```
