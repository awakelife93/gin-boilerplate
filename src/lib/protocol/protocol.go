package protocol

import "github.com/gin-gonic/gin"

type Params struct {
	Message string
	Status  int
	Data    interface{}
}

func Response(c *gin.Context, params Params) {
	c.JSON(params.Status, Params{
		Message: params.Message,
		Status:  params.Status,
		Data:    params.Data,
	})
}

func Request() {
	// todo: write common request logic...
}
