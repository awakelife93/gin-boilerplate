package routes

import (
	"github.com/awakelife93/gin-boilerplate/src/lib"
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	"github.com/gin-gonic/gin"
)

func initialize(engine *gin.Engine) *gin.Engine {

	// todo: add restful apis
	var apis []lib.Api = append(getApis())

	for _, api := range apis {
		engine.Handle(api.Method, api.Path, func(context *gin.Context) {

			item := middleware.GenerateRequestItem(api.Method, context.Request)

			result, error := api.Function(item)

			if error != nil {
				var status int = lib.INTERNAL_SERVER_ERROR_STATUS
				var message string = lib.INTERNAL_SERVER_ERROR_MESSAGE

				if error.Status != nil {
					status = *error.Status
				}

				if error.Message != nil {
					message = *error.Message
				}

				lib.ErrorResponse(context, lib.ErrorResult{
					Message: &message,
					Status:  &status,
					Data:    result,
				})

				return
			}

			lib.Response(context, lib.Result{
				Message: lib.OK_MESSAGE,
				Status:  lib.OK_STATUS,
				Data:    result,
			})
		})
	}

	return engine
}

func Initialize(engine *gin.Engine) *gin.Engine {
	return initialize(engine)
}
