# 代理

Fetch 支持 HTTP、HTTPS 和 SOCKS5 代理。

## 基本代理用法

```go
response, err := fetch.Get("https://httpbin.zcorky.com/ip", &fetch.Config{
	Proxy: "http://127.0.0.1:8080",
})
```

## 代理类型

### HTTP 代理

```go
&fetch.Config{
	Proxy: "http://127.0.0.1:8080",
}
```

### HTTPS 代理

```go
&fetch.Config{
	Proxy: "https://127.0.0.1:8080",
}
```

### SOCKS5 代理

```go
&fetch.Config{
	Proxy: "socks5://127.0.0.1:1080",
}
```

### 带认证的代理

```go
&fetch.Config{
	Proxy: "http://user:password@127.0.0.1:8080",
}
```

## 环境变量

Fetch 自动使用环境变量中的代理设置：

- `HTTP_PROXY`
- `HTTPS_PROXY`
- `SOCKS_PROXY`

如果设置了环境变量，您不需要显式配置它们。

```bash
export HTTP_PROXY=http://127.0.0.1:8080
export HTTPS_PROXY=http://127.0.0.1:8080
```

## 使用 Fetch 实例

```go
f := fetch.New()
f.SetProxy("http://127.0.0.1:8080")

response, err := f.Get("https://httpbin.zcorky.com/ip").Execute()
```
