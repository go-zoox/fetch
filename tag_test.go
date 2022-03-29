package fetch

import "testing"

func TestTagDecode(t *testing.T) {
	type body struct {
		URL      string  `alias:"url"`
		Method   string  `alias:"method"`
		Num      int     `alias:"number"`
		IsBase64 bool    `alias:"is_base64"`
		Float    float64 `alias:"float"`
	}

	var b body
	response := &Response{
		Status: 200,
		Body: []byte(`{
			"url": "/get",
			"method":"GET",
			"number": 10,
			"is_base64": true,
			"float": 1.1
		}`),
	}

	if err := decode(&b, response); err != nil {
		t.Error(err)
	}

	if b.URL != "/get" {
		t.Error("Expected url /get, got", b.URL)
	}

	if b.Method != "GET" {
		t.Error("Expected method GET, got", b.Method)
	}

	if b.Num != 10 {
		t.Error("Expected num 10, got", b.Num)
	}

	if b.IsBase64 != true {
		t.Error("Expected IsBase64 false, got", b.IsBase64)
	}

	if b.Float != 1.1 {
		t.Error("Expected Float 1.1, got", b.Float)
	}
}
