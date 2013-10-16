package main

import (
	"net/http"
	"testing"
)

func TestBaseHandler(t *testing.T) {
	go main()
	time.sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:4321/")
	body := ioutil.ReadAll(resp)
	if body != "{\"message\":\"hello world\"}" {
		t.Error("TestBaseHandler doesn't match:", body)
	}
}
