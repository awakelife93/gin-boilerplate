package lib

import "github.com/gin-gonic/gin"

func CreateEngine() *gin.Engine {
	return gin.Default()
}
