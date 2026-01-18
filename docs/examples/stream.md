# Stream Examples

## Basic Streaming

Stream response data instead of loading it all into memory:

```go
package main

import (
	"io"
	"fmt"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Stream("https://httpbin.zcorky.com/stream/10")
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
		fmt.Print(string(buf[:n]))
	}
}
```

## Stream with Fetch Instance

```go
f := fetch.New()
f.SetURL("https://example.com/stream")
f.SetMethod("GET")

config, _ := f.Config()
config.IsStream = true

response, err := f.SetConfig(config).Execute()
if err != nil {
	panic(err)
}
defer response.Stream.Close()

// Process stream...
```

## Stream Large File Download

```go
package main

import (
	"io"
	"os"
	"github.com/go-zoox/fetch"
)

func main() {
	response, err := fetch.Stream("https://example.com/large-file.zip")
	if err != nil {
		panic(err)
	}
	defer response.Stream.Close()

	file, err := os.Create("/tmp/large-file.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Copy stream to file
	_, err = io.Copy(file, response.Stream)
	if err != nil {
		panic(err)
	}
}
```

## Stream Processing

```go
func processStream(url string, processor func([]byte)) error {
	response, err := fetch.Stream(url)
	if err != nil {
		return err
	}
	defer response.Stream.Close()

	buf := make([]byte, 4096)
	for {
		n, err := response.Stream.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		
		processor(buf[:n])
	}
	
	return nil
}

func main() {
	err := processStream("https://example.com/stream", func(chunk []byte) {
		// Process each chunk
		fmt.Printf("Received %d bytes\n", len(chunk))
	})
	
	if err != nil {
		panic(err)
	}
}
```
