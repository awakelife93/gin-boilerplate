package product

import (
	"github.com/awakelife93/gin-boilerplate/src/lib/protocol"
	service "github.com/awakelife93/gin-boilerplate/src/services/product"
	"github.com/gin-gonic/gin"
)

func GetProducts(context *gin.Context) {
	result := service.GetProducts()

	protocol.Response(context, protocol.Params{
		Message: "hello world",
		Status:  protocol.OK_STATUS,
		Data:    result,
	})
}
