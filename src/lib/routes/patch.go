package routes

import (
	"github.com/awakelife93/gin-boilerplate/src/controllers"
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
)

func patchApis() structures.Apis {
	const method string = "PATCH"
	var apis structures.Apis

	// ************** Product ************** //
	apis = append(
		apis,
		structures.Api{
			Method:   method,
			Path:     "/products/:requestId",
			Function: controllers.UpdateProduct,
		},
	)

	return apis
}
