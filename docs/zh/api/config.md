# Config

`Config` 结构体提供 HTTP 请求的配置选项。

## 类型定义

```go
type Config struct {
	URL     string
	Method  string
	Headers Headers
	Query   Query
	Params  Params
	Body    Body
	BaseURL string
	Timeout time.Duration
	DownloadFilePath string
	Proxy string
	IsStream bool
	IsSession bool
	HTTP2 bool
	
	// TLS 配置
	TLSCaCert     []byte
	TLSCaCertFile string
	TLSCert     []byte
	TLSCertFile string
	TLSKey     []byte
	TLSKeyFile string
	TLSInsecureSkipVerify bool
	
	// Unix 域套接字
	UnixDomainSocket string
	
	// Context
	Context context.Context
	
	// 进度回调
	OnProgress OnProgress
	
	// 认证
	BasicAuth BasicAuth
	Username string
	Password string
}
```

## 字段说明

### 请求配置

- `URL`: 请求的目标 URL
- `Method`: HTTP 方法（GET、POST、PUT、PATCH、DELETE、HEAD）
- `Headers`: 请求头映射
- `Query`: 查询参数映射
- `Params`: URL 路径参数映射
- `Body`: 请求体（可以是 map、string、bytes、io.Reader 等）
- `BaseURL`: 相对路径的基础 URL

### 超时

- `Timeout`: 请求超时时间

### 文件操作

- `DownloadFilePath`: 保存下载文件的路径

### 网络

- `Proxy`: 代理服务器 URL（http、https、socks5）
- `UnixDomainSocket`: Unix 域套接字路径
- `HTTP2`: 启用 HTTP/2 支持

### TLS

- `TLSCaCert`: CA 证书字节
- `TLSCaCertFile`: CA 证书文件路径
- `TLSCert`: 客户端证书字节
- `TLSCertFile`: 客户端证书文件路径
- `TLSKey`: 客户端私钥字节
- `TLSKeyFile`: 客户端私钥文件路径
- `TLSInsecureSkipVerify`: 跳过 TLS 证书验证

### 其他

- `IsStream`: 启用流式模式
- `IsSession`: 启用会话（cookie）管理
- `Context`: 用于取消的 context
- `OnProgress`: 进度回调函数
- `BasicAuth`: 基本认证凭据
- `Username`: 认证用户名
- `Password`: 认证密码

## 方法

### Merge

将另一个配置合并到此配置中。

```go
func (c *Config) Merge(config *Config)
```

### Clone

创建配置的克隆。

```go
func (c *Config) Clone() *Config
```

## 相关类型

### BasicAuth

```go
type BasicAuth struct {
	Username string
	Password string
}
```

### OnProgress

进度回调函数类型。

```go
type OnProgress func(percent int64, current, total int64)
```

## 示例

```go
config := &fetch.Config{
	BaseURL: "https://api.example.com",
	Timeout: 10 * time.Second,
	Headers: map[string]string{
		"Authorization": "Bearer token",
		"Content-Type": "application/json",
	},
	Query: map[string]string{
		"page": "1",
	},
	Body: map[string]interface{}{
		"name": "John",
	},
}

response, err := fetch.Post("/users", config)
```
