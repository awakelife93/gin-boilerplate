package lib

import (
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Method   string
	Path     string
	Function func(middleware.RequestItem) (interface{}, *ErrorResult)
}

type Apis = []Api

type ErrorResult struct {
	Message *string
	Status  *int
	Data    interface{}
}

type Result struct {
	Message string
	Status  int
	Data    interface{}
}

func Response(c *gin.Context, params Result) {
	c.JSON(params.Status, Result{
		Message: params.Message,
		Status:  params.Status,
		Data:    params.Data,
	})
}

func ErrorResponse(c *gin.Context, params ErrorResult) {
	c.JSON(*params.Status, ErrorResult{
		Message: params.Message,
		Status:  params.Status,
		Data:    params.Data,
	})
}

func Request() {
	// todo: write common request logic...
}
