package routes

import (
	cProduct "github.com/awakelife93/gin-boilerplate/src/controllers/product"
)

func getApis() Apis {
	var apis Apis

	// ************** Product ************** //
	apis = append(apis, Api{
		Method:   "GET",
		Path:     "/products",
		Function: cProduct.GetProducts,
	})

	return apis
}
