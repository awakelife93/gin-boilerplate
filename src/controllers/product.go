package controllers

import (
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/awakelife93/gin-boilerplate/src/services"
)

func GetProducts(item structures.RequestItem) (interface{}, *structures.ErrorResult) {
	result, error := services.GetProducts()
	return result, error
}
