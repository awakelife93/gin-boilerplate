package routes

import (
	cProduct "github.com/awakelife93/gin-boilerplate/src/controllers/product"
)

func getApis() Apis {
	var apis Apis

	apis = append(apis, Api{
		Path:     "/products",
		Function: cProduct.GetProducts,
	})

	return apis
}
