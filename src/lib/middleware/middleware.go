package middleware

import (
	"github.com/gin-gonic/gin"
)

func defaultSetupMiddleware(engine *gin.Engine) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
}

func Initialize(engine *gin.Engine) {
	defaultSetupMiddleware(engine)
}
