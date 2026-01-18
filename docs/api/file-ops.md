# File Operations

Functions for downloading, uploading, and streaming files.

## Download

Downloads a file from a URL.

```go
func Download(url string, filepath string, config ...interface{}) (*Response, error)
```

**Parameters:**
- `url`: File URL to download
- `filepath`: Local path to save the file
- `config`: Optional configuration

**Example:**

```go
response, err := fetch.Download("https://example.com/file.zip", "/tmp/file.zip")
```

### Download with Progress

```go
f := fetch.New()
f.SetProgressCallback(func(percent int64, current, total int64) {
	fmt.Printf("Download progress: %d%% (%d/%d bytes)\n", percent, current, total)
})

response, err := f.Download("https://example.com/large-file.zip", "/tmp/file.zip").Execute()
```

## Upload

Uploads a file to a URL.

```go
func Upload(url string, file io.Reader, config ...interface{}) (*Response, error)
```

**Parameters:**
- `url`: Upload endpoint URL
- `file`: File reader (io.Reader)
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

### Upload with Additional Fields

```go
file, err := os.Open("image.jpg")
if err != nil {
	panic(err)
}
defer file.Close()

f := fetch.New()
f.SetMethod("POST")
f.SetURL("https://api.example.com/upload")
f.SetBody(map[string]interface{}{
	"file": file,
	"description": "My image",
	"category": "photos",
})

response, err := f.Execute()
```

### Upload with Progress

```go
file, err := os.Open("large-file.zip")
if err != nil {
	panic(err)
}
defer file.Close()

f := fetch.New()
f.SetProgressCallback(func(percent int64, current, total int64) {
	fmt.Printf("Upload progress: %d%% (%d/%d bytes)\n", percent, current, total)
})

response, err := f.Upload("https://api.example.com/upload", file).Execute()
```

## Stream

Streams response data instead of loading it all into memory.

```go
func Stream(url string, config ...interface{}) (*Response, error)
```

**Parameters:**
- `url`: Request URL
- `config`: Optional configuration

**Example:**

```go
response, err := fetch.Stream("https://example.com/stream")
if err != nil {
	panic(err)
}
defer response.Stream.Close()

// Read stream
buf := make([]byte, 1024)
for {
	n, err := response.Stream.Read(buf)
	if err == io.EOF {
		break
	}
	if err != nil {
		panic(err)
	}
	
	// Process chunk
	processChunk(buf[:n])
}
```

## Using Fetch Instance

All file operations can also be performed using Fetch instances:

```go
f := fetch.New()

// Download
response, err := f.Download("https://example.com/file.zip", "/tmp/file.zip").Execute()

// Upload
file, _ := os.Open("file.txt")
response, err := f.Upload("https://api.example.com/upload", file).Execute()

// Stream
response, err := f.Stream("https://example.com/stream").Execute()
```
