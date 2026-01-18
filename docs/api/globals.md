# Global Functions

Global functions for default configuration and convenience functions.

## Global Configuration

### SetBaseURL

Sets the default base URL for all requests.

```go
func SetBaseURL(url string)
```

**Example:**

```go
fetch.SetBaseURL("https://api.example.com")

// All subsequent requests will use this base URL
response, err := fetch.Get("/users")
```

### SetTimeout

Sets the default timeout for all requests.

```go
func SetTimeout(timeout time.Duration)
```

**Example:**

```go
fetch.SetTimeout(10 * time.Second)

// All subsequent requests will use this timeout
response, err := fetch.Get("https://api.example.com/users")
```

### SetUserAgent

Sets the default user agent for all requests.

```go
func SetUserAgent(userAgent string)
```

**Example:**

```go
fetch.SetUserAgent("MyApp/1.0.0")

// All subsequent requests will use this user agent
response, err := fetch.Get("https://api.example.com/users")
```

## Global Variables

### BaseURL

Default base URL.

```go
var BaseURL string
```

### Timeout

Default timeout.

```go
var Timeout time.Duration
```

### UserAgent

Default user agent.

```go
var UserAgent string
```

## Session

### Session

Creates a new Fetch instance with session enabled (cookies are maintained across requests).

```go
func Session() *Fetch
```

**Example:**

```go
session := fetch.Session()
session.SetBaseURL("https://api.example.com")

// First request - login (sets cookie)
session.Post("/login", &fetch.Config{
	Body: map[string]interface{}{
		"username": "user",
		"password": "pass",
	},
}).Execute()

// Subsequent requests will use cookies from login
response, err := session.Get("/protected").Execute()
```

## Notes

- Global configuration applies to all new Fetch instances created after the configuration is set
- Existing Fetch instances are not affected by global configuration changes
- Use global configuration for convenience, or use instance methods for more control
