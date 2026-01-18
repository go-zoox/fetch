# Config

The `Config` struct provides configuration options for HTTP requests.

## Type Definition

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
	
	// TLS Configuration
	TLSCaCert     []byte
	TLSCaCertFile string
	TLSCert     []byte
	TLSCertFile string
	TLSKey     []byte
	TLSKeyFile string
	TLSInsecureSkipVerify bool
	
	// Unix Domain Socket
	UnixDomainSocket string
	
	// Context
	Context context.Context
	
	// Progress Callback
	OnProgress OnProgress
	
	// Authentication
	BasicAuth BasicAuth
	Username string
	Password string
}
```

## Fields

### Request Configuration

- `URL`: The target URL for the request
- `Method`: HTTP method (GET, POST, PUT, PATCH, DELETE, HEAD)
- `Headers`: Request headers map
- `Query`: Query parameters map
- `Params`: URL path parameters map
- `Body`: Request body (can be map, string, bytes, io.Reader, etc.)
- `BaseURL`: Base URL for relative paths

### Timeout

- `Timeout`: Request timeout duration

### File Operations

- `DownloadFilePath`: Path to save downloaded file

### Network

- `Proxy`: Proxy server URL (http, https, socks5)
- `UnixDomainSocket`: Unix domain socket path
- `HTTP2`: Enable HTTP/2 support

### TLS

- `TLSCaCert`: CA certificate bytes
- `TLSCaCertFile`: Path to CA certificate file
- `TLSCert`: Client certificate bytes
- `TLSCertFile`: Path to client certificate file
- `TLSKey`: Client private key bytes
- `TLSKeyFile`: Path to client private key file
- `TLSInsecureSkipVerify`: Skip TLS certificate verification

### Other

- `IsStream`: Enable streaming mode
- `IsSession`: Enable session (cookie) management
- `Context`: Context for cancellation
- `OnProgress`: Progress callback function
- `BasicAuth`: Basic authentication credentials
- `Username`: Username for authentication
- `Password`: Password for authentication

## Methods

### Merge

Merges another config into this config.

```go
func (c *Config) Merge(config *Config)
```

### Clone

Creates a clone of the config.

```go
func (c *Config) Clone() *Config
```

## Related Types

### BasicAuth

```go
type BasicAuth struct {
	Username string
	Password string
}
```

### OnProgress

Progress callback function type.

```go
type OnProgress func(percent int64, current, total int64)
```

## Example

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
