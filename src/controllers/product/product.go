package product

import (
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	service "github.com/awakelife93/gin-boilerplate/src/services/product"
)

func GetProducts(item middleware.RequestItem) interface{} {
	return service.GetProducts()
}
