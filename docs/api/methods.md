# Methods

Global convenience functions for making HTTP requests.

## Get

Makes a GET request.

```go
func Get(url string, config ...interface{}) (*Response, error)
```

**Parameters:**
- `url`: Request URL
- `config`: Optional configuration

**Example:**

```go
response, err := fetch.Get("https://api.example.com/users")
```

## Post

Makes a POST request.

```go
func Post(url string, config ...interface{}) (*Response, error)
```

**Parameters:**
- `url`: Request URL
- `config`: Optional configuration (should be `*Config`)

**Example:**

```go
response, err := fetch.Post("https://api.example.com/users", &fetch.Config{
	Body: map[string]interface{}{
		"name": "John",
	},
})
```

## Put

Makes a PUT request.

```go
func Put(url string, config ...interface{}) (*Response, error)
```

**Example:**

```go
response, err := fetch.Put("https://api.example.com/users/1", &fetch.Config{
	Body: map[string]interface{}{
		"name": "Jane",
	},
})
```

## Patch

Makes a PATCH request.

```go
func Patch(url string, config ...interface{}) (*Response, error)
```

**Example:**

```go
response, err := fetch.Patch("https://api.example.com/users/1", &fetch.Config{
	Body: map[string]interface{}{
		"email": "new@example.com",
	},
})
```

## Delete

Makes a DELETE request.

```go
func Delete(url string, config ...interface{}) (*Response, error)
```

**Example:**

```go
response, err := fetch.Delete("https://api.example.com/users/1")
```

## Head

Makes a HEAD request.

```go
func Head(url string, config ...interface{}) (*Response, error)
```

**Example:**

```go
response, err := fetch.Head("https://api.example.com/users/1")
```

## Download

Downloads a file from a URL.

```go
func Download(url string, filepath string, config ...interface{}) (*Response, error)
```

**Parameters:**
- `url`: File URL
- `filepath`: Path to save the file
- `config`: Optional configuration

**Example:**

```go
response, err := fetch.Download("https://example.com/file.zip", "/tmp/file.zip")
```

## Upload

Uploads a file.

```go
func Upload(url string, file io.Reader, config ...interface{}) (*Response, error)
```

**Parameters:**
- `url`: Upload endpoint URL
- `file`: File reader
- `config`: Optional configuration

**Example:**

```go
file, err := os.Open("local-file.txt")
if err != nil {
	panic(err)
}
defer file.Close()

response, err := fetch.Upload("https://api.example.com/upload", file)
```
