# Response

`Response` 类型表示 HTTP 响应。

## 类型定义

```go
type Response struct {
	Status  int
	Headers http.Header
	Body    []byte
	Request *Config
	Stream  io.ReadCloser
}
```

## 字段

- `Status`: HTTP 状态码
- `Headers`: HTTP 响应头
- `Body`: 响应体字节
- `Request`: 原始请求配置
- `Stream`: 响应流（当 IsStream 为 true 时）

## 方法

### String

将响应体作为字符串返回。

```go
func (r *Response) String() string
```

### Value

将响应体作为 gjson.Result 返回，用于 JSON 解析。

```go
func (r *Response) Value() gjson.Result
```

### Get

通过键路径获取 JSON 值。

```go
func (r *Response) Get(key string) gjson.Result
```

**示例：**

```go
value := response.Get("user.name")
array := response.Get("items.#")
```

### JSON

将响应体作为格式化的 JSON 字符串返回。

```go
func (r *Response) JSON() (string, error)
```

### UnmarshalJSON

将响应体反序列化为 JSON 结构体。

```go
func (r *Response) UnmarshalJSON(v interface{}) error
```

**示例：**

```go
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var user User
err := response.UnmarshalJSON(&user)
```

### UnmarshalYAML

将响应体反序列化为 YAML 结构体。

```go
func (r *Response) UnmarshalYAML(v interface{}) error
```

### Ok

如果状态码为 2xx 则返回 true。

```go
func (r *Response) Ok() bool
```

### Error

返回包含状态码和响应体的错误。

```go
func (r *Response) Error() error
```

### StatusCode

返回 HTTP 状态码。

```go
func (r *Response) StatusCode() int
```

### StatusText

返回 HTTP 状态文本。

```go
func (r *Response) StatusText() string
```

### ContentType

返回 Content-Type 请求头。

```go
func (r *Response) ContentType() string
```

### Location

返回 Location 请求头。

```go
func (r *Response) Location() string
```

### ContentLength

返回 Content-Length 请求头的值。

```go
func (r *Response) ContentLength() int
```

### ContentEncoding

返回 Content-Encoding 请求头。

```go
func (r *Response) ContentEncoding() string
```

### TransferEncoding

返回 Transfer-Encoding 请求头。

```go
func (r *Response) TransferEncoding() string
```

### ContentLanguage

返回 Content-Language 请求头。

```go
func (r *Response) ContentLanguage() string
```

### XPoweredBy

返回 X-Powered-By 请求头。

```go
func (r *Response) XPoweredBy() string
```

### XRequestID

返回 X-Request-ID 请求头。

```go
func (r *Response) XRequestID() string
```

### AcceptRanges

返回 Accept-Ranges 请求头。

```go
func (r *Response) AcceptRanges() string
```

### SetCookie

返回 Set-Cookie 请求头。

```go
func (r *Response) SetCookie() string
```

## 示例

```go
response, err := fetch.Get("https://api.example.com/users/1")
if err != nil {
	panic(err)
}

if response.Ok() {
	// 解析 JSON
	user := response.Get("user")
	fmt.Println(user.Get("name"))
	
	// 或反序列化为结构体
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	var userObj User
	err := response.UnmarshalJSON(&userObj)
	fmt.Println(userObj.Name)
} else {
	fmt.Printf("错误: %v\n", response.Error())
}
```
