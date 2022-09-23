package routes

import (
	"github.com/awakelife93/gin-boilerplate/src/lib"
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/gin-gonic/gin"
)

func generateGroupByApis(engine *gin.Engine) {
	// * you can write APIs for each group here.
	// * sample group & api
	v1 := engine.Group("/v1")

	v1.GET("/sample", func(context *gin.Context) {
		lib.Response(context, structures.Result{
			Message: lib.OK_MESSAGE,
			Status:  lib.OK_STATUS,
			Data:    nil,
		})
	})
}

func initialize(engine *gin.Engine) *gin.Engine {

	// * if you declare an api for each group
	generateGroupByApis(engine)

	apis := append(getApis())

	for _, api := range apis {
		// * protected overwritten api
		thisApi := api

		engine.Handle(thisApi.Method, thisApi.Path, func(context *gin.Context) {

			item := middleware.GenerateRequestItem(thisApi.Method, context.Request)

			result, error := thisApi.Function(item)

			if error != nil {
				status := lib.INTERNAL_SERVER_ERROR_STATUS
				message := lib.INTERNAL_SERVER_ERROR_MESSAGE

				if error.Status != nil {
					status = *error.Status
				}

				if error.Message != nil {
					message = *error.Message
				}

				lib.ErrorResponse(context, structures.ErrorResult{
					Message: &message,
					Status:  &status,
					Data:    result,
				})

				return
			}

			lib.Response(context, structures.Result{
				Message: lib.OK_MESSAGE,
				Status:  lib.OK_STATUS,
				Data:    result,
			})
		})
	}

	return engine
}

func Initialize(engine *gin.Engine) {
	initialize(engine)
}
