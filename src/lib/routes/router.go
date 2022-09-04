package routes

import (
	"github.com/gin-gonic/gin"
)

type Api struct {
	Path     string
	Function func(*gin.Context)
}

type Apis = []Api

func setupRouters(engine *gin.Engine) *gin.Engine {

	getApiLists := getApis()
	for i := 0; i < len(getApiLists); i++ {
		engine.GET(getApiLists[i].Path, getApiLists[i].Function)
	}

	return engine
}

func Initialize(engine *gin.Engine) *gin.Engine {
	return setupRouters(engine)
}
