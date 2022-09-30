package controllers

import (
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/awakelife93/gin-boilerplate/src/services"
	"github.com/awakelife93/gin-boilerplate/src/utils"
)

func GetProducts(item structures.RequestItem) (interface{}, *structures.ErrorResult) {
	result, error := services.GetProducts()
	return result, error
}

func GetProduct(item structures.RequestItem) (interface{}, *structures.ErrorResult) {
	id := utils.GetRequestId(item.Item)
	result, error := services.GetProduct(id)
	return result, error
}
