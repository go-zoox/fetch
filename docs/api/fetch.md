# Fetch

The `Fetch` type is the main client for making HTTP requests.

## Functions

### New

Creates a new fetch client.

```go
func New(cfg ...*Config) *Fetch
```

**Parameters:**
- `cfg`: Optional configuration

**Example:**

```go
f := fetch.New()
f := fetch.New(&fetch.Config{
	Timeout: 10 * time.Second,
})
```

### Create

Creates a new fetch with base URL.

```go
func Create(baseURL string) *Fetch
```

**Parameters:**
- `baseURL`: Base URL for all requests

**Example:**

```go
f := fetch.Create("https://api.example.com")
```

## Methods

### SetURL

Sets the request URL.

```go
func (f *Fetch) SetURL(url string) *Fetch
```

### SetBaseURL

Sets the base URL.

```go
func (f *Fetch) SetBaseURL(url string) *Fetch
```

### SetMethod

Sets the HTTP method.

```go
func (f *Fetch) SetMethod(method string) *Fetch
```

### SetHeader

Sets a request header.

```go
func (f *Fetch) SetHeader(key, value string) *Fetch
```

### SetQuery

Sets a query parameter.

```go
func (f *Fetch) SetQuery(key, value string) *Fetch
```

### SetBody

Sets the request body.

```go
func (f *Fetch) SetBody(body Body) *Fetch
```

### SetTimeout

Sets the request timeout.

```go
func (f *Fetch) SetTimeout(timeout time.Duration) *Fetch
```

### SetBasicAuth

Sets basic authentication.

```go
func (f *Fetch) SetBasicAuth(username, password string) *Fetch
```

### SetBearerToken

Sets bearer token authentication.

```go
func (f *Fetch) SetBearerToken(token string) *Fetch
```

### SetProxy

Sets proxy server.

```go
func (f *Fetch) SetProxy(proxy string) *Fetch
```

### SetContext

Sets the context for cancellation.

```go
func (f *Fetch) SetContext(ctx context.Context) *Fetch
```

### SetProgressCallback

Sets progress callback for upload/download.

```go
func (f *Fetch) SetProgressCallback(callback func(percent int64, current, total int64)) *Fetch
```

### Get

Sets HTTP method to GET.

```go
func (f *Fetch) Get(url string, config ...*Config) *Fetch
```

### Post

Sets HTTP method to POST.

```go
func (f *Fetch) Post(url string, config ...*Config) *Fetch
```

### Put

Sets HTTP method to PUT.

```go
func (f *Fetch) Put(url string, config ...*Config) *Fetch
```

### Patch

Sets HTTP method to PATCH.

```go
func (f *Fetch) Patch(url string, config ...*Config) *Fetch
```

### Delete

Sets HTTP method to DELETE.

```go
func (f *Fetch) Delete(url string, config ...*Config) *Fetch
```

### Head

Sets HTTP method to HEAD.

```go
func (f *Fetch) Head(url string, config ...*Config) *Fetch
```

### Execute

Executes the request.

```go
func (f *Fetch) Execute() (*Response, error)
```

### Send

Alias for Execute.

```go
func (f *Fetch) Send() (*Response, error)
```

### Retry

Retries the request with optional modifications.

```go
func (f *Fetch) Retry(before func(f *Fetch)) (*Response, error)
```

### Clone

Creates a clone of the fetch instance.

```go
func (f *Fetch) Clone() *Fetch
```

### Config

Returns the built configuration.

```go
func (f *Fetch) Config() (*Config, error)
```

## Example

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
