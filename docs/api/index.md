# API Reference

This section documents all public APIs of the fetch package.

## Overview

The fetch package provides:

- **Global Functions**: Convenient functions for making HTTP requests (`Get`, `Post`, `Put`, `Patch`, `Delete`, `Head`, `Download`, `Upload`, `Stream`)
- **Global Configuration**: Functions to set default configuration (`SetBaseURL`, `SetTimeout`, `SetUserAgent`)
- **Fetch Type**: Main client type with chainable methods
- **Config Type**: Configuration for HTTP requests
- **Response Type**: HTTP response handler
- **Session**: Session management for maintaining cookies across requests

## Quick Links

- [Global Functions](./globals) - Global configuration and convenience functions
- [Fetch](./fetch) - Main client type and instance methods
- [Config](./config) - Request configuration
- [Response](./response) - Response handling
- [Methods](./methods) - HTTP method functions (Get, Post, Put, Patch, Delete, Head)
- [File Operations](./file-ops) - Download, Upload, and Stream functions

## Global Functions

### HTTP Methods

```go
// GET request
response, err := fetch.Get(url, config)

// POST request
response, err := fetch.Post(url, config)

// PUT request
response, err := fetch.Put(url, config)

// PATCH request
response, err := fetch.Patch(url, config)

// DELETE request
response, err := fetch.Delete(url, config)

// HEAD request
response, err := fetch.Head(url, config)
```

### File Operations

```go
// Download file
response, err := fetch.Download(url, filepath, config)

// Upload file
response, err := fetch.Upload(url, file, config)

// Stream response
response, err := fetch.Stream(url, config)
```

### Session

```go
// Create a session (maintains cookies)
session := fetch.Session()
```

## Main Types

### Fetch

The main client type for making HTTP requests. Supports method chaining.

```go
f := fetch.New()
f.SetBaseURL("https://api.example.com")
f.SetBearerToken("token")
response, err := f.Get("/users").Execute()
```

See [Fetch API](./fetch) for details.

### Config

Configuration struct for customizing HTTP requests.

```go
config := &fetch.Config{
	Timeout: 10 * time.Second,
	Headers: map[string]string{
		"Authorization": "Bearer token",
	},
}
```

See [Config API](./config) for details.

### Response

Response handler with JSON parsing capabilities.

```go
response, err := fetch.Get(url)
if err != nil {
	panic(err)
}

json := response.JSON()
value := response.Get("key")
```

See [Response API](./response) for details.
