package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-zoox/headers"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"
)

// Response is the fetch response
type Response struct {
	Status      int
	Headers     http.Header
	Body        []byte
	resultCache gjson.Result
	parsed      bool
	//
	Request *Config
	//
	Stream io.ReadCloser
}

// String returns the body as string
func (r *Response) String() string {
	return string(r.Body)
}

// Value returns the body as gjson.Result
func (r *Response) Value() gjson.Result {
	if !r.parsed {
		r.resultCache = gjson.Parse(r.String())
		r.parsed = true
	}

	return r.resultCache
}

// Get returns the value of the key
func (r *Response) Get(key string) gjson.Result {
	return r.Value().Get(key)
}

// JSON returns the body as json string
func (r *Response) JSON() (string, error) {
	raw := r.String()
	b, err := json.MarshalIndent(gjson.Parse(raw).Value(), "", "  ")
	if err != nil {
		return "", errors.New("invalid json: " + raw)
	}

	return string(b), nil
}

// func (r *Response) Unmarshal(v interface{}) error {
// 	return json.Unmarshal(r.Body, v)
// 	// return decode(v, r)
// }

// UnmarshalJSON unmarshals body to json struct
//
// @TODO bug when lint (go vet) method UnmarshalJSON(v interface{}) error should have signature UnmarshalJSON([]byte) error
func (r *Response) UnmarshalJSON(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}

// UnmarshalYAML unmarshals body to yaml struct
func (r *Response) UnmarshalYAML(v interface{}) error {
	return yaml.Unmarshal(r.Body, v)
}

// Ok returns true if status code is 2xx
func (r *Response) Ok() bool {
	return r.Status >= 200 && r.Status < 300
}

// Error returns error with status and response string.
func (r *Response) Error() error {
	return fmt.Errorf("[%d] %s", r.Status, r.String())
}

// StatusCode returns status code of the response
func (r *Response) StatusCode() int {
	return r.Status
}

// StatusText returns status text of the response
func (r *Response) StatusText() string {
	return http.StatusText(r.Status)
}

// ContentType returns content type of the response
func (r *Response) ContentType() string {
	return r.Headers.Get(headers.ContentType)
}

// Location returns location of the response
func (r *Response) Location() string {
	return r.Headers.Get(headers.Location)
}

// ContentLength returns content length of the response
func (r *Response) ContentLength() int {
	vs := r.Headers.Get(headers.ContentLength)
	if vs == "" {
		return 0
	}

	value, err := strconv.Atoi(vs)
	if err != nil {
		return 0
	}

	return value
}

// ContentEncoding returns content encoding of the response
func (r *Response) ContentEncoding() string {
	return r.Headers.Get(headers.ContentEncoding)
}

// TransferEncoding returns transfer encoding of the response
func (r *Response) TransferEncoding() string {
	return r.Headers.Get(headers.TransferEncoding)
}

// ContentLanguage returns content language of the response
func (r *Response) ContentLanguage() string {
	return r.Headers.Get(headers.ContentLanguage)
}

// XPoweredBy returns x-powered-by of the response
func (r *Response) XPoweredBy() string {
	return r.Headers.Get(headers.XPoweredBy)
}

// XRequestID returns x-request-id of the response
func (r *Response) XRequestID() string {
	return r.Headers.Get(headers.XRequestID)
}

// AcceptRanges returns x-accept-ranges of the response
func (r *Response) AcceptRanges() string {
	return r.Headers.Get(headers.AcceptRanges)
}

// SetCookie returns set-cookie of the response
func (r *Response) SetCookie() string {
	return r.Headers.Get(headers.SetCookie)
}
