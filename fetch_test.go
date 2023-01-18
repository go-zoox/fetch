package fetch

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-zoox/headers"
	"github.com/go-zoox/testify"
)

func TestNew(t *testing.T) {
	f := New()
	// testify.NotNil(t, f)
	testify.Assert(t, f != nil, "Expected not nil, got nil")
}

func TestSetMethod(t *testing.T) {
	f := New()
	f.SetMethod("GET")
	testify.Equal(t, "GET", f.config.Method)

	// invalid method
	f.SetMethod("INVALID")
	testify.Assert(t, &f.Errors[0] != nil)
}

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
		// SetProxy("https://127.0.0.1:17890").
		// SetProxy("socks5://127.0.0.1:17890").
		Send()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("response:", response.String())
}

func TestRetryManual(t *testing.T) {
	f := New()

	response, err := f.Get("https://httpbin.zcorky.com/headers").
		SetBearToken("zzz").
		Send()
	if err != nil {
		t.Fatal(err)
	}

	// j, _ := json.MarshalIndent(f.config, " ", "	")
	// fmt.Println(string(j))

	if response.Get("headers.authorization").String() != "Bearer zzz" {
		t.Fatal("Expected Authorization Bearer zzz, got", response.Get("headers.authorization").String())
	}

	response, err = f.Retry(func(f *Fetch) {
		f.SetBearToken("another")
	})
	if err != nil {
		t.Fatal(err)
	}

	if response.Get("headers.authorization").String() != "Bearer another" {
		t.Fatal("Expected Authorization Bearer zzz, got", response.Get("headers.authorization").String())
	}
}

func TestCreate(t *testing.T) {
	baseURL := "https://httpbin.zcorky.com"
	f := Create(baseURL)
	testify.Equal(t, baseURL, f.config.BaseURL)
}

func TestSetQuery(t *testing.T) {
	f := New()
	f.SetQuery("a", "b")
	f.SetQuery("c", "d")

	testify.Equal(t, "b", f.config.Query.Get("a"))
	testify.Equal(t, "d", f.config.Query.Get("c"))
}

func TestSetHeader(t *testing.T) {
	f := New()
	f.SetHeader("a", "b")
	f.SetHeader("c", "d")

	testify.Equal(t, "b", f.config.Headers.Get("a"))
	testify.Equal(t, "d", f.config.Headers.Get("c"))
}

func TestSetParams(t *testing.T) {
	f := New()
	f.SetParam("a", "b")
	f.SetParam("c", "d")

	testify.Equal(t, "b", f.config.Params.Get("a"))
	testify.Equal(t, "d", f.config.Params.Get("c"))
}

func TestSetBody(t *testing.T) {
	f := New()
	f.SetBody("a")
	testify.Equal(t, "a", f.config.Body.(string))

	// body := map[string]string{
	// 	"a": "b",
	// }
	// f.SetBody(body)
	// testify.Equal(t, body, f.config.Body.(map[string]string))
}

func TestSetBaseURL(t *testing.T) {
	f := New()
	f.SetBaseURL("https://httpbin.zcorky.com")
	testify.Equal(t, "https://httpbin.zcorky.com", f.config.BaseURL)
}

func TestSetTimeout(t *testing.T) {
	f := New()
	f.SetTimeout(1 * time.Second)
	testify.Equal(t, 1*time.Second, f.config.Timeout)
}

func TestSetUserAgent(t *testing.T) {
	f := New()
	f.SetUserAgent("test")
	testify.Equal(t, "test", f.config.Headers.Get(headers.UserAgent))
}

func TestConfigSetBasicAuth(t *testing.T) {
	f := New()
	f.SetBasicAuth("user", "passwd")
	testify.Equal(t, "Basic dXNlcjpwYXNzd2Q", f.config.Headers.Get(headers.Authorization))
}

func TestConfigSetBearToken(t *testing.T) {
	f := New()
	f.SetBearToken("token")
	testify.Equal(t, "Bearer token", f.config.Headers.Get(headers.Authorization))
}

func TestSetAuthorization(t *testing.T) {
	f := New()
	f.SetAuthorization("token")
	testify.Equal(t, "token", f.config.Headers.Get(headers.Authorization))
}

func TestSetAccept(t *testing.T) {
	f := New()
	f.SetAccept("application/json")
	testify.Equal(t, "application/json", f.config.Headers.Get(headers.Accept))
}

func TestSetContentType(t *testing.T) {
	f := New()
	f.SetContentType("application/json")
	testify.Equal(t, "application/json", f.config.Headers.Get(headers.ContentType))
}

func TestSetReferrer(t *testing.T) {
	f := New()
	f.SetReferrer("https://httpbin.zcorky.com")
	testify.Equal(t, "https://httpbin.zcorky.com", f.config.Headers.Get(headers.Referrer))
}

func TestSetCacheControl(t *testing.T) {
	f := New()
	f.SetCacheControl("no-cache")
	testify.Equal(t, "no-cache", f.config.Headers.Get(headers.CacheControl))
}

func TestSetAcceptEncoding(t *testing.T) {
	f := New()
	f.SetAcceptEncoding("gzip")
	testify.Equal(t, "gzip", f.config.Headers.Get(headers.AcceptEncoding))
}

func TestSetAcceptLanguage(t *testing.T) {
	f := New()
	f.SetAcceptLanguage("zh-CN")
	testify.Equal(t, "zh-CN", f.config.Headers.Get(headers.AcceptLanguage))
}

func TestSetProxy(t *testing.T) {
	f := New()
	f.SetProxy("https://example.com")
	testify.Equal(t, "https://example.com", f.config.Proxy)
}

func TestFetchConfig(t *testing.T) {
	f := New()
	f.SetBaseURL("https://httpbin.zcorky.com")
	f.SetURL("/get/:id/:name")
	f.SetParam("id", "1")
	f.SetParam("name", "Zero")
	cfg, err := f.Config()
	testify.Assert(t, err == nil, "err should be nil")
	testify.Equal(t, "https://httpbin.zcorky.com/get/1/Zero", cfg.URL)
}

func TestFetchCancel(t *testing.T) {
	f := New()
	f.SetBaseURL("https://httpbin.zcorky.com")
	f.SetURL("/delay/3")
	ctx, cancel := context.WithCancel(context.Background())
	f.SetContext(ctx)
	cancel()
	_, err := f.Execute()
	// fmt.Println(err)
	testify.Assert(t, err != nil, "err should not be nil")
}
