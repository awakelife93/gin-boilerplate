package routes

import (
	"github.com/awakelife93/gin-boilerplate/src/controllers"
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
)

func getApis() structures.Apis {
	const method string = "GET"
	var apis structures.Apis

	// ************** Product ************** //
	apis = append(apis, structures.Api{
		Method:   method,
		Path:     "/products",
		Function: controllers.GetProducts,
	})

	return apis
}
