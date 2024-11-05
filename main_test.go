package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	homeHelloWord(res, req, nil)

	if res.Code != http.StatusOK {
		t.Errorf("Expected 200 as status code in GET to /, received %d", res.Code)
	}
}
