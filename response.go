package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-yaml"
	"github.com/tidwall/gjson"
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

func (r *Response) Error() error {
	return fmt.Errorf("[%d] %s", r.Status, r.String())
}
