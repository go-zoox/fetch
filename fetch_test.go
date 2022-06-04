package fetch

import (
	"errors"
	"fmt"
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
		t.Fatal("Expected BaseURL https://httpbin.zcorky.com, got", response.Get("origin").String())
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
	if err := response.UnmarshalJSON(&b); err != nil {
		t.Error(err)
	}

	if b.URL != "/get" {
		t.Error("Expected url /get, got", b.URL)
	}

	if b.Method != "GET" {
		t.Error("Expected method GET, got", b.Method)
	}
}

func TestSetBasicAuth(t *testing.T) {
	f := New()

	response, err := f.Get("https://httpbin.zcorky.com/basic-auth/user/passwd").
		SetBasicAuth("user", "passwd").
		Send()
	if err != nil {
		t.Error(err)
	}

	if response.Status != 200 {
		t.Error("Expected authenticated 200, got", response.Status)
	}
}

func TestSetBearToken(t *testing.T) {
	f := New()

	response, err := f.Get("https://httpbin.zcorky.com/headers").
		SetBearToken("token").
		Send()
	if err != nil {
		t.Error(err)
	}

	if response.Status != 200 {
		t.Error("Expected authenticated 200, got", response.Status)
	}

	if response.Get("headers.authorization").String() != "Bearer token" {
		t.Error("Expected Authorization Bearer token, got", response.Get("headers.authorization").String())
	}
}

func TestProxy(t *testing.T) {
	f := New()

	response, err := f.Get("https://httpbin.org/ip").
		// SetProxy("http://127.0.0.1:17890").
		Send()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("response:", response.String())
}
