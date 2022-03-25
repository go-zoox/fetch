package fetch

import (
	"errors"
	"testing"
	"time"
)

func TestBaseURL(t *testing.T) {
	BaseURL := "https://httpbin.zcorky.com"

	f := New()

	response, err := f.Get("/get", &Config{BaseURL: BaseURL}).Send()
	if err != nil {
		t.Error(err)
	}

	if response.Get("origin").String() != BaseURL {
		t.Error("Expected BaseURL https://httpbin.zcorky.com, got", response.Get("origin").String())
	}
}

func TestTimeout(t *testing.T) {
	BaseURL := "https://httpbin.zcorky.com"

	f := New()

	_, err := f.Get("/get", &Config{
		BaseURL: BaseURL,
		Timeout: 1 * time.Microsecond,
	}).Send()
	if err == nil {
		t.Error(errors.New("Expected timeout error, got nil"))
	}
}
