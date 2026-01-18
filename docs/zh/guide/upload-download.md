# 上传与下载

## 下载文件

从 URL 下载文件：

```go
response, err := fetch.Download("https://example.com/image.jpg", "/tmp/image.jpg")
if err != nil {
	panic(err)
}
```

### 带进度的下载

```go
f := fetch.New()
f.SetProgressCallback(func(percent int64, current, total int64) {
	fmt.Printf("下载进度: %d%% (%d/%d 字节)\n", percent, current, total)
})

response, err := f.Download("https://example.com/large-file.zip", "/tmp/file.zip").Execute()
```

## 上传文件

上传文件：

```go
file, err := os.Open("local-file.txt")
if err != nil {
	panic(err)
}
defer file.Close()

response, err := fetch.Upload("https://api.example.com/upload", file)
if err != nil {
	panic(err)
}

fmt.Println(response.JSON())
```

### 带额外表单字段的上传

```go
file, err := os.Open("image.jpg")
if err != nil {
	panic(err)
}
defer file.Close()

f := fetch.New()
f.SetMethod("POST")
f.SetURL("https://api.example.com/upload")
f.SetBody(map[string]interface{}{
	"file": file,
	"description": "My image",
	"category": "photos",
})

response, err := f.Execute()
```

## 示例：完整的文件传输

```go
package main

import (
	"fmt"
	"os"
	"github.com/go-zoox/fetch"
)

func main() {
	// 下载
	fmt.Println("正在下载文件...")
	_, err := fetch.Download("https://example.com/file.zip", "/tmp/file.zip")
	if err != nil {
		panic(err)
	}
	fmt.Println("下载完成！")

	// 上传
	fmt.Println("正在上传文件...")
	file, err := os.Open("/tmp/file.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	response, err := fetch.Upload("https://api.example.com/upload", file)
	if err != nil {
		panic(err)
	}
	fmt.Println("上传完成！", response.JSON())
}
```
