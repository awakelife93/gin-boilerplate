package middleware

import (
	"github.com/awakelife93/gin-boilerplate/src/config"
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/awakelife93/gin-boilerplate/src/utils"
	"github.com/gin-gonic/gin"
)

func GenerateRequestItem(method string, context *gin.Context) structures.RequestItem {
	if method == "GET" || method == "DELETE" {

		return structures.RequestItem{
			Item: utils.GenerateQueryItem(
				context.Request.URL.Query(),
				context.Params.ByName(config.RequestIdFieldName()),
			),
		}
	}

	return structures.RequestItem{
		Item: utils.GenerateBodyItem(
			context.Request.Body,
		),
	}
}
