# Fetch

`Fetch` 类型是用于发起 HTTP 请求的主客户端。

## 函数

### New

创建一个新的 fetch 客户端。

```go
func New(cfg ...*Config) *Fetch
```

**参数：**
- `cfg`: 可选配置

**示例：**

```go
f := fetch.New()
f := fetch.New(&fetch.Config{
	Timeout: 10 * time.Second,
})
```

### Create

创建一个带有基础 URL 的新 fetch 实例。

```go
func Create(baseURL string) *Fetch
```

**参数：**
- `baseURL`: 所有请求的基础 URL

**示例：**

```go
f := fetch.Create("https://api.example.com")
```

## 方法

### SetURL

设置请求 URL。

```go
func (f *Fetch) SetURL(url string) *Fetch
```

### SetBaseURL

设置基础 URL。

```go
func (f *Fetch) SetBaseURL(url string) *Fetch
```

### SetMethod

设置 HTTP 方法。

```go
func (f *Fetch) SetMethod(method string) *Fetch
```

### SetHeader

设置请求头。

```go
func (f *Fetch) SetHeader(key, value string) *Fetch
```

### SetQuery

设置查询参数。

```go
func (f *Fetch) SetQuery(key, value string) *Fetch
```

### SetBody

设置请求体。

```go
func (f *Fetch) SetBody(body Body) *Fetch
```

### SetTimeout

设置请求超时。

```go
func (f *Fetch) SetTimeout(timeout time.Duration) *Fetch
```

### SetBasicAuth

设置基本认证。

```go
func (f *Fetch) SetBasicAuth(username, password string) *Fetch
```

### SetBearerToken

设置 Bearer Token 认证。

```go
func (f *Fetch) SetBearerToken(token string) *Fetch
```

### SetProxy

设置代理服务器。

```go
func (f *Fetch) SetProxy(proxy string) *Fetch
```

### SetContext

设置用于取消的 context。

```go
func (f *Fetch) SetContext(ctx context.Context) *Fetch
```

### SetProgressCallback

设置上传/下载的进度回调。

```go
func (f *Fetch) SetProgressCallback(callback func(percent int64, current, total int64)) *Fetch
```

### Get

设置 HTTP 方法为 GET。

```go
func (f *Fetch) Get(url string, config ...*Config) *Fetch
```

### Post

设置 HTTP 方法为 POST。

```go
func (f *Fetch) Post(url string, config ...*Config) *Fetch
```

### Put

设置 HTTP 方法为 PUT。

```go
func (f *Fetch) Put(url string, config ...*Config) *Fetch
```

### Patch

设置 HTTP 方法为 PATCH。

```go
func (f *Fetch) Patch(url string, config ...*Config) *Fetch
```

### Delete

设置 HTTP 方法为 DELETE。

```go
func (f *Fetch) Delete(url string, config ...*Config) *Fetch
```

### Head

设置 HTTP 方法为 HEAD。

```go
func (f *Fetch) Head(url string, config ...*Config) *Fetch
```

### Execute

执行请求。

```go
func (f *Fetch) Execute() (*Response, error)
```

### Send

Execute 的别名。

```go
func (f *Fetch) Send() (*Response, error)
```

### Retry

重试请求，可选修改。

```go
func (f *Fetch) Retry(before func(f *Fetch)) (*Response, error)
```

### Clone

创建 fetch 实例的克隆。

```go
func (f *Fetch) Clone() *Fetch
```

### Config

返回构建的配置。

```go
func (f *Fetch) Config() (*Config, error)
```

## 示例

```go
f := fetch.New()
f.SetBaseURL("https://api.example.com")
f.SetBearerToken("token")
f.SetTimeout(10 * time.Second)

response, err := f.Get("/users").Execute()
if err != nil {
	panic(err)
}

fmt.Println(response.JSON())
```
