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

func TestResponseUnmarshal(t *testing.T) {
	type body struct {
		URL    string `alias:"url"`
		Method string `alias:"method"`
	}

	var b body
	response, _ := Get("https://httpbin.zcorky.com/get")
	if err := response.Unmarshal(&b); err != nil {
		t.Error(err)
	}

	if b.URL != "/get" {
		t.Error("Expected url /get, got", b.URL)
	}

	if b.Method != "GET" {
		t.Error("Expected method GET, got", b.Method)
	}
}
