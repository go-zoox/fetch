# 文件操作

用于下载、上传和流式传输文件的函数。

## Download

从 URL 下载文件。

```go
func Download(url string, filepath string, config ...interface{}) (*Response, error)
```

**参数：**
- `url`: 要下载的文件 URL
- `filepath`: 保存文件的本地路径
- `config`: 可选配置

**示例：**

```go
response, err := fetch.Download("https://example.com/file.zip", "/tmp/file.zip")
```

### 带进度的下载

```go
f := fetch.New()
f.SetProgressCallback(func(percent int64, current, total int64) {
	fmt.Printf("下载进度: %d%% (%d/%d 字节)\n", percent, current, total)
})

response, err := f.Download("https://example.com/large-file.zip", "/tmp/file.zip").Execute()
```

## Upload

将文件上传到 URL。

```go
func Upload(url string, file io.Reader, config ...interface{}) (*Response, error)
```

**参数：**
- `url`: 上传端点 URL
- `file`: 文件读取器（io.Reader）
- `config`: 可选配置

**示例：**

```go
file, err := os.Open("local-file.txt")
if err != nil {
	panic(err)
}
defer file.Close()

response, err := fetch.Upload("https://api.example.com/upload", file)
```

### 带额外字段的上传

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

### 带进度的上传

```go
file, err := os.Open("large-file.zip")
if err != nil {
	panic(err)
}
defer file.Close()

f := fetch.New()
f.SetProgressCallback(func(percent int64, current, total int64) {
	fmt.Printf("上传进度: %d%% (%d/%d 字节)\n", percent, current, total)
})

response, err := f.Upload("https://api.example.com/upload", file).Execute()
```

## Stream

流式传输响应数据，而不是将其全部加载到内存中。

```go
func Stream(url string, config ...interface{}) (*Response, error)
```

**参数：**
- `url`: 请求 URL
- `config`: 可选配置

**示例：**

```go
response, err := fetch.Stream("https://example.com/stream")
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
	processChunk(buf[:n])
}
```

## 使用 Fetch 实例

所有文件操作也可以使用 Fetch 实例执行：

```go
f := fetch.New()

// 下载
response, err := f.Download("https://example.com/file.zip", "/tmp/file.zip").Execute()

// 上传
file, _ := os.Open("file.txt")
response, err := f.Upload("https://api.example.com/upload", file).Execute()

// 流式传输
response, err := f.Stream("https://example.com/stream").Execute()
```
