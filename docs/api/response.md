# Response

The `Response` type represents an HTTP response.

## Type Definition

```go
type Response struct {
	Status  int
	Headers http.Header
	Body    []byte
	Request *Config
	Stream  io.ReadCloser
}
```

## Fields

- `Status`: HTTP status code
- `Headers`: HTTP response headers
- `Body`: Response body bytes
- `Request`: Original request configuration
- `Stream`: Response stream (when IsStream is true)

## Methods

### String

Returns the response body as a string.

```go
func (r *Response) String() string
```

### Value

Returns the response body as a gjson.Result for JSON parsing.

```go
func (r *Response) Value() gjson.Result
```

### Get

Gets a JSON value by key path.

```go
func (r *Response) Get(key string) gjson.Result
```

**Example:**

```go
value := response.Get("user.name")
array := response.Get("items.#")
```

### JSON

Returns the response body as formatted JSON string.

```go
func (r *Response) JSON() (string, error)
```

### UnmarshalJSON

Unmarshals the response body into a JSON struct.

```go
func (r *Response) UnmarshalJSON(v interface{}) error
```

**Example:**

```go
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var user User
err := response.UnmarshalJSON(&user)
```

### UnmarshalYAML

Unmarshals the response body into a YAML struct.

```go
func (r *Response) UnmarshalYAML(v interface{}) error
```

### Ok

Returns true if status code is 2xx.

```go
func (r *Response) Ok() bool
```

### Error

Returns an error with status code and response body.

```go
func (r *Response) Error() error
```

### StatusCode

Returns the HTTP status code.

```go
func (r *Response) StatusCode() int
```

### StatusText

Returns the HTTP status text.

```go
func (r *Response) StatusText() string
```

### ContentType

Returns the Content-Type header.

```go
func (r *Response) ContentType() string
```

### Location

Returns the Location header.

```go
func (r *Response) Location() string
```

### ContentLength

Returns the Content-Length header value.

```go
func (r *Response) ContentLength() int
```

### ContentEncoding

Returns the Content-Encoding header.

```go
func (r *Response) ContentEncoding() string
```

### TransferEncoding

Returns the Transfer-Encoding header.

```go
func (r *Response) TransferEncoding() string
```

### ContentLanguage

Returns the Content-Language header.

```go
func (r *Response) ContentLanguage() string
```

### XPoweredBy

Returns the X-Powered-By header.

```go
func (r *Response) XPoweredBy() string
```

### XRequestID

Returns the X-Request-ID header.

```go
func (r *Response) XRequestID() string
```

### AcceptRanges

Returns the Accept-Ranges header.

```go
func (r *Response) AcceptRanges() string
```

### SetCookie

Returns the Set-Cookie header.

```go
func (r *Response) SetCookie() string
```

## Example

```go
response, err := fetch.Get("https://api.example.com/users/1")
if err != nil {
	panic(err)
}

if response.Ok() {
	// Parse JSON
	user := response.Get("user")
	fmt.Println(user.Get("name"))
	
	// Or unmarshal to struct
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	var userObj User
	err := response.UnmarshalJSON(&userObj)
	fmt.Println(userObj.Name)
} else {
	fmt.Printf("Error: %v\n", response.Error())
}
```
