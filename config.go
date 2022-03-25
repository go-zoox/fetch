package fetch

import "strings"

type Config struct {
	Url     string
	Method  string
	Headers Headers
	Query   Query
	Params  Params
	Body    Body
	//
	BaseURL string
}

type Body interface{}

type Headers map[string]string

func (h Headers) Get(key string) string {
	for k, v := range h {
		if strings.ToLower(k) == strings.ToLower(key) {
			return strings.ToLower(v)
		}
	}

	return ""
}

type Query map[string]string

func (h Query) Get(key string) string {
	return h[key]
}

type Params map[string]string

func (h Params) Get(key string) string {
	return h[key]
}
