package routes

import (
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	"github.com/awakelife93/gin-boilerplate/src/lib/protocol"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Method   string
	Path     string
	Function func(middleware.RequestItem) interface{}
}

type Apis = []Api

func initialize(engine *gin.Engine) *gin.Engine {

	// todo: add restful apis
	var apis []Api = append(getApis())

	for _, api := range apis {
		engine.Handle(api.Method, api.Path, func(context *gin.Context) {
			// * generate common item
			item := middleware.GenerateRequestItem(api.Method, context.Request)

			// * api result
			result := api.Function(item)

			protocol.Response(context, protocol.Params{
				Message: protocol.OK_MESSAGE,
				Status:  protocol.OK_STATUS,
				Data:    result,
			})
		})
	}

	return engine
}

func Initialize(engine *gin.Engine) *gin.Engine {
	return initialize(engine)
}
