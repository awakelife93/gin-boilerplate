package lib

import "github.com/gin-gonic/gin"

func CreateEngine(mode string) *gin.Engine {
	gin.SetMode(mode)
	return gin.Default()
}
