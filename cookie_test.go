package fetch

import (
	"testing"
)

func TestCookieEmpty(t *testing.T) {
	cookie := newCookie()
	if cookie.Get("") != "" {
		t.Error("Expected empty string, got", cookie.Get(""))
	}

	if err := cookie.Add("", ""); err == nil {
		t.Fatal("Expected error, got nil")
	}

	if items := cookie.Items(); len(items) != 0 {
		t.Fatal("Expected 0 items, got", len(items))
	}

	if err := cookie.Set("key", "value"); err != nil {
		t.Fatal("Expected error, got nil")
	}

	if items := cookie.Items(); len(items) != 1 {
		t.Fatal("Expected 1 items, got", len(items))
	}

	if cookie.Get("key") != "value" {
		t.Fatal("Expected value, got", cookie.Get("key"))
	}

	if cookie.String() != "key=value" {
		t.Fatal("Expected key=value, got", cookie.String())
	}

	if err := cookie.Set("key1", "value1"); err != nil {
		t.Fatal("Expected error, got nil")
	}

	if items := cookie.Items(); len(items) != 2 {
		t.Fatal("Expected 2 items, got", len(items))
	}

	if cookie.Get("key1") != "value1" {
		t.Fatal("Expected value, got", cookie.Get("key1"))
	}

	if cookie.String() != "key=value; key1=value1" {
		t.Fatal("Expected string, got", cookie.String())
	}

	if err := cookie.Remove("key"); err != nil {
		t.Fatal("Expected error, got nil")
	}

	if cookie.Get("key") != "" {
		t.Fatal("Expected empty string, got", cookie.Get("key1"))
	}

	if cookie.Get("key1") != "value1" {
		t.Fatal("Expected value, got", cookie.Get("key1"))
	}

	if cookie.String() != "key1=value1" {
		t.Fatal("Expected key1=value1, got", cookie.String())
	}

	if err := cookie.Clear(); err != nil {
		t.Fatal("Expected error, got nil")
	}

	if cookie.Get("key1") != "" {
		t.Fatal("Expected empty string, got", cookie.Get("key1"))
	}

	if cookie.String() != "" {
		t.Fatal("Expected empty string, got", cookie.String())
	}
}

func TestCookieParse(t *testing.T) {
	cookie := newCookie()
	if err := cookie.Parse("key=value; key1=value1"); err != nil {
		t.Fatal("Expected error, got nil")
	}

	if cookie.Get("key") != "value" {
		t.Fatal("Expected value, got", cookie.Get("key"))
	}

	if cookie.Get("key1") != "value1" {
		t.Fatal("Expected value, got", cookie.Get("key1"))
	}

	if cookie.String() != "key=value; key1=value1" {
		t.Fatal("Expected string, got", cookie.String())
	}
}

func TestCookie(t *testing.T) {
	response, err := New().
		SetCookie("key", "value").
		SetCookie("key1", "value1").
		SetURL("https://httpbin.zcorky.com/headers").
		Send()
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(response.Get("headers").String())

	if response.Get("headers.cookie").String() != "key=value; key1=value1" {
		t.Fatal("Expected cookie, got", response.Get("header.cookie").String())
	}
}
