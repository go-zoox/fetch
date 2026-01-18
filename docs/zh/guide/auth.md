# 认证

Fetch 支持多种认证方法。

## Basic 认证

```go
response, err := fetch.Get("https://httpbin.zcorky.com/basic-auth/user/pass", &fetch.Config{
	BasicAuth: &fetch.BasicAuth{
		Username: "user",
		Password: "pass",
	},
})
```

### 使用 Fetch 实例

```go
f := fetch.New()
f.SetBasicAuth("user", "pass")

response, err := f.Get("https://httpbin.zcorky.com/basic-auth/user/pass").Execute()
```

## Bearer Token

```go
response, err := fetch.Get("https://api.example.com/protected", &fetch.Config{
	Headers: map[string]string{
		"Authorization": "Bearer your-token-here",
	},
})
```

### 使用 SetBearerToken

```go
f := fetch.New()
f.SetBearerToken("your-token-here")

response, err := f.Get("https://api.example.com/protected").Execute()
```

## 自定义认证请求头

```go
response, err := fetch.Get("https://api.example.com/protected", &fetch.Config{
	Headers: map[string]string{
		"Authorization": "Custom token-here",
	},
})
```
