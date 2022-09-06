package middleware

import (
	"net/http"

	"github.com/awakelife93/gin-boilerplate/src/utils"
)

func GenerateRequestItem(method string, request *http.Request) RequestItem {
	if method == "GET" || method == "DELETE" {
		return RequestItem{item: utils.GenerateQueryItem(request.URL.Query())}
	}

	return RequestItem{item: utils.GenerateBodyItem(request.Body)}
}
