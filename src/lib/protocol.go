package lib

import (
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, params structures.Result) {
	c.JSON(params.Status, structures.Result{
		Message: params.Message,
		Status:  params.Status,
		Data:    params.Data,
	})
}

func ErrorResponse(c *gin.Context, params structures.ErrorResult) {
	c.JSON(*params.Status, structures.ErrorResult{
		Message: params.Message,
		Status:  params.Status,
		Data:    params.Data,
	})
}

func Request() {
	// todo: write common request logic...
}
