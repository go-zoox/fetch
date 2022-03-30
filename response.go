package fetch

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/tidwall/gjson"
)

type Response struct {
	Status      int
	Headers     http.Header
	Body        []byte
	resultCache gjson.Result
	parsed      bool
}

func (r *Response) String() string {
	return string(r.Body)
}

func (r *Response) Value() gjson.Result {
	if !r.parsed {
		r.resultCache = gjson.Parse(r.String())
		r.parsed = true
	}

	return r.resultCache
}

func (r *Response) Get(key string) gjson.Result {
	return r.Value().Get(key)
}

func (r *Response) JSON() (string, error) {
	raw := r.String()
	b, err := json.MarshalIndent(gjson.Parse(raw).Value(), "", "  ")
	if err != nil {
		return "", errors.New("invalid json: " + raw)
	}

	return string(b), nil
}

func (r *Response) Unmarshal(v interface{}) error {
	return json.Unmarshal(r.Body, v)
	// return decode(v, r)
}
