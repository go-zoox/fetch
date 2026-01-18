# Methods

用于发起 HTTP 请求的全局便捷函数。

## Get

发起 GET 请求。

```go
func Get(url string, config ...interface{}) (*Response, error)
```

**参数：**
- `url`: 请求 URL
- `config`: 可选配置

**示例：**

```go
response, err := fetch.Get("https://api.example.com/users")
```

## Post

发起 POST 请求。

```go
func Post(url string, config ...interface{}) (*Response, error)
```

**参数：**
- `url`: 请求 URL
- `config`: 可选配置（应为 `*Config`）

**示例：**

```go
response, err := fetch.Post("https://api.example.com/users", &fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
	},
})
```

## Put

发起 PUT 请求。

```go
func Put(url string, config ...interface{}) (*Response, error)
```

**示例：**

```go
response, err := fetch.Put("https://api.example.com/users/1", &fetch.Config{
	Body: map[string]interface{}{
		"name": "Jane",
	},
})
```

## Patch

发起 PATCH 请求。

```go
func Patch(url string, config ...interface{}) (*Response, error)
```

**示例：**

```go
response, err := fetch.Patch("https://api.example.com/users/1", &fetch.Config{
	Body: map[string]interface{}{
		"email": "new@example.com",
	},
})
```

## Delete

发起 DELETE 请求。

```go
func Delete(url string, config ...interface{}) (*Response, error)
```

**示例：**

```go
response, err := fetch.Delete("https://api.example.com/users/1")
```

## Head

发起 HEAD 请求。

```go
func Head(url string, config ...interface{}) (*Response, error)
```

**示例：**

```go
response, err := fetch.Head("https://api.example.com/users/1")
```

## Download

从 URL 下载文件。

```go
func Download(url string, filepath string, config ...interface{}) (*Response, error)
```

**参数：**
- `url`: 文件 URL
- `filepath`: 保存文件的路径
- `config`: 可选配置

**示例：**

```go
response, err := fetch.Download("https://example.com/file.zip", "/tmp/file.zip")
```

## Upload

上传文件。

```go
func Upload(url string, file io.Reader, config ...interface{}) (*Response, error)
```

**参数：**
- `url`: 上传端点 URL
- `file`: 文件读取器
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
