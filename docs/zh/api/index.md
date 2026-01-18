# API 参考

本节文档介绍了 fetch 包的所有公共 API。

## 概述

fetch 包提供：

- **全局函数**：用于发起 HTTP 请求的便捷函数（`Get`、`Post`、`Put`、`Patch`、`Delete`、`Head`、`Download`、`Upload`、`Stream`）
- **全局配置**：用于设置默认配置的函数（`SetBaseURL`、`SetTimeout`、`SetUserAgent`）
- **Fetch 类型**：具有链式方法的主客户端类型
- **Config 类型**：HTTP 请求的配置
- **Response 类型**：HTTP 响应处理器
- **Session**：用于在请求之间维护 cookie 的会话管理

## 快速链接

- [全局函数](./globals) - 全局配置和便捷函数
- [Fetch](./fetch) - 主客户端类型和实例方法
- [Config](./config) - 请求配置
- [Response](./response) - 响应处理
- [Methods](./methods) - HTTP 方法函数（Get、Post、Put、Patch、Delete、Head）
- [文件操作](./file-ops) - Download、Upload 和 Stream 函数

## 全局函数

### HTTP 方法

```go
// GET 请求
response, err := fetch.Get(url, config)

// POST 请求
response, err := fetch.Post(url, config)

// PUT 请求
response, err := fetch.Put(url, config)

// PATCH 请求
response, err := fetch.Patch(url, config)

// DELETE 请求
response, err := fetch.Delete(url, config)

// HEAD 请求
response, err := fetch.Head(url, config)

// 下载文件
response, err := fetch.Download(url, filepath, config)

// 上传文件
response, err := fetch.Upload(url, file, config)
```

## 主要类型

### Fetch

用于发起 HTTP 请求的主客户端类型。支持方法链式调用。

```go
f := fetch.New()
f.SetBaseURL("https://api.example.com")
f.SetBearerToken("token")
response, err := f.Get("/users").Execute()
```

详见 [Fetch API](./fetch)。

### Config

用于自定义 HTTP 请求的配置结构体。

```go
config := &fetch.Config{
	Timeout: 10 * time.Second,
	Headers: map[string]string{
		"Authorization": "Bearer token",
	},
}
```

详见 [Config API](./config)。

### Response

具有 JSON 解析功能的响应处理器。

```go
response, err := fetch.Get(url)
if err != nil {
	panic(err)
}

json := response.JSON()
value := response.Get("key")
```

详见 [Response API](./response)。
