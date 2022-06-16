package fetch

import (
	"fmt"
	"testing"
	"time"
)

func TestGlobalSetBaseURL(t *testing.T) {
	if BaseURL != "" {
		t.Errorf("global BaseURL should be empty string, but got %s", BaseURL)
	}

	SetBaseURL("https://example.com")
	if BaseURL != "https://example.com" {
		t.Errorf("global BaseURL should be https://example.com, but got %s", BaseURL)
	}

	SetBaseURL("")
	if BaseURL != "" {
		t.Errorf("global BaseURL should be empty string, but got %s", BaseURL)
	}
}

func TestGlobalUserAgent(t *testing.T) {
	if UserAgent != fmt.Sprintf("GoFetch/%s (github.com/go-zoox/fetch)", Version) {
		t.Errorf("global UserAgent should be empty string, but got %s", UserAgent)
	}

	SetUserAgent("test user agent")
	if UserAgent != "test user agent" {
		t.Errorf("global UserAgent should be test user agent, but got %s", UserAgent)
	}

	SetUserAgent(fmt.Sprintf("GoFetch/%s (github.com/go-zoox/fetch)", Version))
	if UserAgent != fmt.Sprintf("GoFetch/%s (github.com/go-zoox/fetch)", Version) {
		t.Errorf("global UserAgent should be empty string, but got %s", UserAgent)
	}
}

func TestGlobalTimeout(t *testing.T) {
	if Timeout != 60*time.Second {
		t.Errorf("global Timeout should be empty string, but got %s", Timeout)
	}

	SetTimeout(30 * time.Second)
	if Timeout != 30*time.Second {
		t.Errorf("global Timeout should be test user agent, but got %s", Timeout)
	}

	SetTimeout(60 * time.Second)
	if Timeout != 60*time.Second {
		t.Errorf("global Timeout should be empty string, but got %s", Timeout)
	}
}
