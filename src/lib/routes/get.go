package routes

import (
	"github.com/awakelife93/gin-boilerplate/src/controllers"
	"github.com/awakelife93/gin-boilerplate/src/lib"
)

func getApis() lib.Apis {
	const method string = "GET"
	var apis lib.Apis

	// ************** Product ************** //
	apis = append(apis, lib.Api{
		Method:   method,
		Path:     "/products",
		Function: controllers.GetProducts,
	})

	// ************** Node ************** //
	apis = append(apis, lib.Api{
		Method:   method,
		Path:     "/nodes",
		Function: controllers.GetNodes,
	})

	return apis
}
