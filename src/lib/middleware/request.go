package middleware

import (
	"net/http"

	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/awakelife93/gin-boilerplate/src/utils"
)

func GenerateRequestItem(method string, request *http.Request) structures.RequestItem {
	if method == "GET" || method == "DELETE" {
		return structures.RequestItem{
			Item: utils.GenerateQueryItem(
				request.URL.Query(),
			),
		}
	}

	return structures.RequestItem{
		Item: utils.GenerateBodyItem(
			request.Body,
		),
	}
}
