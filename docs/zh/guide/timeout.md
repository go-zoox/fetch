# 超时与重试

## 超时

设置请求超时以防止请求挂起：

```go
response, err := fetch.Get("https://httpbin.zcorky.com/get", &fetch.Config{
	Timeout: 5 * time.Second,
})
```

### 使用 Fetch 实例

```go
f := fetch.New()
f.SetTimeout(10 * time.Second)

response, err := f.Get("https://httpbin.zcorky.com/get").Execute()
```

### 使用 Context 设置超时

您也可以使用 Go 的 context 来获得更多控制：

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

f := fetch.New()
f.SetContext(ctx)

response, err := f.Get("https://httpbin.zcorky.com/get").Execute()
```

## 重试

使用 `Retry` 方法重试失败的请求：

```go
f := fetch.New()
f.SetURL("https://httpbin.zcorky.com/get")

response, err := f.Retry(func(nf *fetch.Fetch) {
	// 如果需要，在重试前修改请求
	nf.SetHeader("X-Retry-Count", "1")
})
```

## 最佳实践

1. **始终设置超时**以防止请求挂起
2. **使用 context 取消**以获得更细粒度的控制
3. **对临时性失败实施指数退避重试**
4. **记录重试尝试**以便调试
