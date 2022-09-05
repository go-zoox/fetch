package fetch

import (
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	file, _ := os.Open("go.mod")

	response, err := Upload("https://httpbin.zcorky.com/upload", file)
	if err != nil {
		t.Error(err)
	}

	if response.Status != 200 {
		t.Error("Expected status code 200, got", response.Status)
	}

	if response.Get("files.file.name").String() != "go.mod" {
		t.Error("Expected file go.mod, got", response.Get("files.file.name").String())
	}
}
