package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	req.RemoteAddr = "1.2.3.4"
	res := httptest.NewRecorder()
	HandleHello(res, req)

	if res.Body.String() != "Hello, 1.2.3.4" {
		t.Error("Wrong /hello response")
	}
}
