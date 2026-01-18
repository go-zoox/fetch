# 配置

`Config` 结构体提供了各种选项来自定义 HTTP 请求。

## 基本配置

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	Timeout: 10 * time.Second,
	Headers: map[string]string{
		"User-Agent": "MyApp/1.0",
	},
})
```

## 配置选项

### URL 和方法

- `URL`: 请求的目标 URL
- `Method`: HTTP 方法（GET、POST、PUT、PATCH、DELETE、HEAD）
- `BaseURL`: 相对路径的基础 URL

### 请求头

```go
&fetch.Config{
	Headers: map[string]string{
		"Authorization": "Bearer token",
		"Content-Type": "application/json",
	},
}
```

### 查询参数

```go
&fetch.Config{
	Query: map[string]string{
		"page": "1",
		"limit": "10",
	},
}
```

### 请求体

```go
&fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
		"age": 30,
	},
}
```

### 超时

```go
&fetch.Config{
	Timeout: 5 * time.Second,
}
```

### 代理

```go
&fetch.Config{
	Proxy: "http://127.0.0.1:8080",
}
```

### 认证

```go
// Basic 认证
&fetch.Config{
	BasicAuth: &fetch.BasicAuth{
		Username: "user",
		Password: "pass",
	},
}

// 或通过请求头使用 Bearer Token
&fetch.Config{
	Headers: map[string]string{
		"Authorization": "Bearer token",
	},
}
```

### TLS 配置

```go
&fetch.Config{
	TLSCaCertFile: "/path/to/ca.crt",
	TLSInsecureSkipVerify: false,
}
```

### Unix 域套接字

```go
&fetch.Config{
	UnixDomainSocket: "/var/run/docker.sock",
}
```

## 使用 Fetch 实例

您也可以配置一个 Fetch 实例并重复使用：

```go
f := fetch.New()
f.SetBaseURL("https://api.example.com")
f.SetTimeout(10 * time.Second)
f.SetHeader("Authorization", "Bearer token")

response, err := f.Get("/users").Execute()
```
