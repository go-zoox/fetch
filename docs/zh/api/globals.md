# 全局函数

用于默认配置和便捷函数的全局函数。

## 全局配置

### SetBaseURL

设置所有请求的默认基础 URL。

```go
func SetBaseURL(url string)
```

**示例：**

```go
fetch.SetBaseURL("https://api.example.com")

// 所有后续请求将使用此基础 URL
response, err := fetch.Get("/users")
```

### SetTimeout

设置所有请求的默认超时。

```go
func SetTimeout(timeout time.Duration)
```

**示例：**

```go
fetch.SetTimeout(10 * time.Second)

// 所有后续请求将使用此超时
response, err := fetch.Get("https://api.example.com/users")
```

### SetUserAgent

设置所有请求的默认 User-Agent。

```go
func SetUserAgent(userAgent string)
```

**示例：**

```go
fetch.SetUserAgent("MyApp/1.0.0")

// 所有后续请求将使用此 User-Agent
response, err := fetch.Get("https://api.example.com/users")
```

## 全局变量

### BaseURL

默认基础 URL。

```go
var BaseURL string
```

### Timeout

默认超时。

```go
var Timeout time.Duration
```

### UserAgent

默认 User-Agent。

```go
var UserAgent string
```

## Session

### Session

创建一个启用 session 的新 Fetch 实例（在请求之间维护 cookie）。

```go
func Session() *Fetch
```

**示例：**

```go
session := fetch.Session()
session.SetBaseURL("https://api.example.com")

// 第一个请求 - 登录（设置 cookie）
session.Post("/login", &fetch.Config{
	Body: map[string]interface{}{
		"username": "user",
		"password": "pass",
	},
}).Execute()

// 后续请求将使用登录时的 cookie
response, err := session.Get("/protected").Execute()
```

## 注意事项

- 全局配置仅适用于配置设置后创建的新 Fetch 实例
- 现有的 Fetch 实例不受全局配置更改的影响
- 使用全局配置便于使用，或使用实例方法以获得更多控制
