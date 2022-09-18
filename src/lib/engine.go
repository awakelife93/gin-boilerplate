package lib

import "github.com/gin-gonic/gin"

func initializeMode(mode string) {
	if mode == gin.DebugMode {
		gin.ForceConsoleColor()
	} else {
		gin.DisableConsoleColor()
	}

	gin.SetMode(mode)
}

func CreateEngine(mode string) *gin.Engine {
	initializeMode(mode)
	return gin.Default()
}
