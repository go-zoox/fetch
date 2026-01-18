# 高级特性

## HTTP/2 支持

启用 HTTP/2：

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	HTTP2: true,
})
```

## TLS 配置

### 自定义 CA 证书

```go
response, err := fetch.Get("https://internal-api.example.com", &fetch.Config{
	TLSCaCertFile: "/path/to/ca.crt",
})
```

或从字节读取：

```go
caCert, _ := os.ReadFile("/path/to/ca.crt")

response, err := fetch.Get("https://internal-api.example.com", &fetch.Config{
	TLSCaCert: caCert,
})
```

### 客户端证书认证

```go
response, err := fetch.Get("https://api.example.com", &fetch.Config{
	TLSCertFile: "/path/to/client.crt",
	TLSKeyFile:  "/path/to/client.key",
})
```

### 跳过 TLS 验证（不推荐用于生产环境）

```go
response, err := fetch.Get("https://self-signed.example.com", &fetch.Config{
	TLSInsecureSkipVerify: true,
})
```

## Unix 域套接字

通过 Unix 域套接字连接：

```go
response, err := fetch.Get("http://localhost/info", &fetch.Config{
	UnixDomainSocket: "/var/run/docker.sock",
})
```

## 流式传输

流式传输响应数据：

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

## Context 取消

使用 context 取消请求：

```go
ctx, cancel := context.WithCancel(context.Background())

f := fetch.New()
f.SetContext(ctx)
f.SetURL("https://slow-api.example.com")

// 在 goroutine 中启动请求
go func() {
	response, err := f.Execute()
	if err != nil {
		fmt.Println("请求已取消或失败:", err)
		return
	}
	fmt.Println(response.JSON())
}()

// 1 秒后取消
time.Sleep(1 * time.Second)
cancel()
```
