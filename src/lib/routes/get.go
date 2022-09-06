package routes

import (
	"github.com/awakelife93/gin-boilerplate/src/controllers"
	"github.com/awakelife93/gin-boilerplate/src/lib"
)

func getApis() lib.Apis {
	var apis lib.Apis

	// ************** Product ************** //
	apis = append(apis, lib.Api{
		Method:   "GET",
		Path:     "/products",
		Function: controllers.GetProducts,
	})

	// ************** Node ************** //
	apis = append(apis, lib.Api{
		Method:   "GET",
		Path:     "/nodes",
		Function: controllers.GetProducts,
	})

	return apis
}
