package assignment2

import (
	"net/http"
	"testing"
)

func TestGetQueryLimit(t *testing.T) {
	client := http.DefaultClient
	r := GetTheRequest("http://google.com?limit=5", client)
	limit, err := GetQueryLimit(r.Request)
	if err != nil {
		t.Errorf("Query works")
	}
	if limit != 5 {
		t.Errorf("Query dont work")
	}
}

func TestAuth(t *testing.T) {
	client := http.DefaultClient
	r := GetTheRequest("http://google.com?auth=testauth", client)
	auth := Auth(r.Request)
	if auth != "testauth" {
		t.Errorf("Query dont work")
	}
}
