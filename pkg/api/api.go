package api

import (
	"fmt"
	"net/http"
)

func HandleHello(res http.ResponseWriter, req *http.Request) {
	var replyType string

	reply, ok := req.URL.Query()["reply"]
	if ok && len(reply) > 0 {
		replyType = reply[0]
	}

	switch replyType {
	case "not-found":
		http.NotFound(res, req)
	case "err":
		http.Error(res, "error", 500)
	default:
		fmt.Fprintf(res, "Hello, %v", req.RemoteAddr)
	}
}
