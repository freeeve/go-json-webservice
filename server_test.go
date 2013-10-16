package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestBaseHandler(t *testing.T) {
	go main()

	time.Sleep(200 * time.Millisecond)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Get("https://localhost:4321/")
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "{\"message\":\"hello world\"}" {
		t.Error("TestBaseHandler doesn't match:", body)
	}
}
