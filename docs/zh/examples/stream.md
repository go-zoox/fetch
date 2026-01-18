# 流式传输示例

## 基本流式传输

流式传输响应数据，而不是将其全部加载到内存中：

```go
package main

import (
	"io"
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Stream("https://httpbin.zcorky.com/stream/10")
	if err != nil {
		panic(err)
	}
	defer response.Stream.Close()

	// 读取流
	buf := make([]byte, 1024)
	for {
		n, err := response.Stream.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		
		// 处理数据块
		fmt.Print(string(buf[:n]))
	}
}
```

## 使用 Fetch 实例流式传输

```go
f := fetch.New()
f.SetURL("https://example.com/stream")
f.SetMethod("GET")

config, _ := f.Config()
config.IsStream = true

response, err := f.SetConfig(config).Execute()
if err != nil {
	panic(err)
}
defer response.Stream.Close()

// 处理流...
```

## 流式下载大文件

```go
package main

import (
	"io"
	"os"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Stream("https://example.com/large-file.zip")
	if err != nil {
		panic(err)
	}
	defer response.Stream.Close()

	file, err := os.Create("/tmp/large-file.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 将流复制到文件
	_, err = io.Copy(file, response.Stream)
	if err != nil {
		panic(err)
	}
}
```
