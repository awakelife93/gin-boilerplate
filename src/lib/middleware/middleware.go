package middleware

import "github.com/gin-gonic/gin"

func defaultSetupMiddleware(engine *gin.Engine) *gin.Engine {
	engine.Use(gin.Logger())

	engine.Use(gin.Recovery())

	return engine
}

func Initialize(engine *gin.Engine) *gin.Engine {
	engine = defaultSetupMiddleware(engine)

	// todo: add middleware

	return engine
}
