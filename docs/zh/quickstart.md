# 快速开始

## 安装

使用 Go 模块安装包：

```bash
go get github.com/go-zoox/fetch
```

## 第一个请求

下面是一个简单的示例：

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Get("https://httpbin.zcorky.com/get")
	if err != nil {
		panic(err)
	}

	url := response.Get("url")
	method := response.Get("method")

	fmt.Println(url, method)
}
```

## 基本 GET 请求

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Get("https://httpbin.zcorky.com/get")
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```

## 基本 POST 请求

```go
package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Post("https://httpbin.zcorky.com/post", &fetch.Config{
		Body: map[string]interface{}{
			"foo": "bar",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.JSON())
}
```

## 下一步

- 阅读[指南](/zh/guide/)了解详细文档
- 查看[示例](/zh/examples/)了解更多用例
- 浏览[API 参考](/zh/api/)查看完整 API 文档
