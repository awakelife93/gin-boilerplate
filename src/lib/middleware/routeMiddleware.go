package middleware

import (
	"net/http"
)

type RequestItem struct {
	item interface{}
}

func GenerateRequestItem(method string, request *http.Request) RequestItem {
	if method == "GET" || method == "DELETE" {
		return RequestItem{item: request.URL.Query()}
	}

	return RequestItem{item: request.Body}
}
