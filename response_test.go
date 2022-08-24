package fetch

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/go-zoox/testify"
)

func TestResponse(t *testing.T) {
	os.Setenv(EnvDEBUG, "true")

	r := &Response{}
	testify.Equal(t, r.Status, 0)

	r.Body, _ = json.Marshal(map[string]string{
		"key": "value",
	})
	testify.Equal(t, string(r.Body), `{"key":"value"}`)

	testify.Equal(t, r.String(), `{"key":"value"}`)

	jsonStr, _ := r.JSON()
	testify.Equal(t, "{\n  \"key\": \"value\"\n}", jsonStr)

	r.Status = 400
	testify.Equal(t, "[400] {\"key\":\"value\"}", r.Error().Error())

	testify.Equal(t, "Bad Request", r.StatusText())

	// content type
	testify.Equal(t, "", r.ContentType())
	r.Headers = http.Header{}
	r.Headers.Set("content-type", "text/plain; charset=utf-8")
	testify.Equal(t, "text/plain; charset=utf-8", r.ContentType())

	// location
	testify.Equal(t, "", r.Location())
	r.Headers.Set("location", "https://httpbin.zcorky.com/image")
	testify.Equal(t, "https://httpbin.zcorky.com/image", r.Location())

	// content length
	testify.Equal(t, 0, r.ContentLength())
	r.Headers.Set("content-length", "10")
	testify.Equal(t, 10, r.ContentLength())

	// transfer encoding
	testify.Equal(t, "", r.TransferEncoding())
	r.Headers.Set("transfer-encoding", "chunked")
	testify.Equal(t, "chunked", r.TransferEncoding())

	// content language
	testify.Equal(t, "", r.ContentLanguage())
	r.Headers.Set("content-language", "en-US")
	testify.Equal(t, "en-US", r.ContentLanguage())

	// content encoding
	testify.Equal(t, "", r.ContentEncoding())
	r.Headers.Set("content-encoding", "gzip")
	testify.Equal(t, "gzip", r.ContentEncoding())

	// x-powered-by
	testify.Equal(t, "", r.XPoweredBy())
	r.Headers.Set("x-powered-by", "Go")
	testify.Equal(t, "Go", r.XPoweredBy())

	// x-request-id
	testify.Equal(t, "", r.XRequestID())
	r.Headers.Set("x-request-id", "12345")
	testify.Equal(t, "12345", r.XRequestID())

	// accept-ranges
	testify.Equal(t, "", r.AcceptRanges())
	r.Headers.Set("accept-ranges", "bytes")
	testify.Equal(t, "bytes", r.AcceptRanges())

	// set cookie
	testify.Equal(t, "", r.SetCookie())
	r.Headers.Set("set-cookie", "key=value")
	testify.Equal(t, "key=value", r.SetCookie())

	// ok
	r.Status = 200
	testify.Equal(t, true, r.Ok())
	r.Status = 400
	testify.Equal(t, false, r.Ok())
	r.Status = 500
	testify.Equal(t, false, r.Ok())
	r.Status = 201
	testify.Equal(t, true, r.Ok())
	r.Status = 202
	testify.Equal(t, true, r.Ok())

	// status code
	testify.Equal(t, 202, r.StatusCode())

	// DEBUG
	os.Setenv(EnvDEBUG, "")
}
