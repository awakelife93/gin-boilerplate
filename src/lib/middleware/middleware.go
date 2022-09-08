package middleware

import (
	"github.com/gin-gonic/gin"
)

type RequestItem struct {
	item interface{}
}

func defaultSetupMiddleware(engine *gin.Engine) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
}

func Initialize(engine *gin.Engine) *gin.Engine {
	defaultSetupMiddleware(engine)
	return engine
}
