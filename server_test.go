package main

import (
	"net/http"
	"testing"
   "time"
   "io/ioutil"
)

func TestBaseHandler(t *testing.T) {
	go main()
	time.Sleep(100 * time.Millisecond)

	resp, _ := http.Get("http://localhost:4321/")
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "{\"message\":\"hello world\"}" {
		t.Error("TestBaseHandler doesn't match:", body)
	}
}
