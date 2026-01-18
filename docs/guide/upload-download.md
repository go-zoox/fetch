# Upload & Download

## Download Files

Download a file from a URL:

```go
response, err := fetch.Download("https://example.com/image.jpg", "/tmp/image.jpg")
if err != nil {
	panic(err)
}
```

### Download with Progress

```go
f := fetch.New()
f.SetProgressCallback(func(percent int64, current, total int64) {
	fmt.Printf("Download progress: %d%% (%d/%d bytes)\n", percent, current, total)
})

response, err := f.Download("https://example.com/large-file.zip", "/tmp/file.zip").Execute()
```

## Upload Files

Upload a file:

```go
file, err := os.Open("local-file.txt")
if err != nil {
	panic(err)
}
defer file.Close()

response, err := fetch.Upload("https://api.example.com/upload", file)
if err != nil {
	panic(err)
}

fmt.Println(response.JSON())
```

### Upload with Additional Form Fields

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

## Example: Complete File Transfer

```go
package main

import (
	"fmt"
	"os"
	"github.com/go-zoox/fetch"
)

func main() {
	// Download
	fmt.Println("Downloading file...")
	_, err := fetch.Download("https://example.com/file.zip", "/tmp/file.zip")
	if err != nil {
		panic(err)
	}
	fmt.Println("Download complete!")

	// Upload
	fmt.Println("Uploading file...")
	file, err := os.Open("/tmp/file.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	response, err := fetch.Upload("https://api.example.com/upload", file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Upload complete!", response.JSON())
}
```
