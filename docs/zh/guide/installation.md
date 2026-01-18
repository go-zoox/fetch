# 安装

## 要求

- Go 1.18 或更高版本

## 安装

使用 Go 模块安装包：

```bash
go get github.com/go-zoox/fetch
```

## 导入

在 Go 代码中导入包：

```go
import "github.com/go-zoox/fetch"
```

## 验证安装

您可以通过创建一个简单的测试文件来验证安装：

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
	fmt.Println("安装成功！", response.Status)
}
```

运行它：

```bash
go run main.go
```

如果您看到"安装成功！"和状态码，说明一切正常！
