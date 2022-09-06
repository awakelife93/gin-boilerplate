package controllers

import (
	"github.com/awakelife93/gin-boilerplate/src/lib"
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	"github.com/awakelife93/gin-boilerplate/src/services"
)

func GetProducts(item middleware.RequestItem) (interface{}, *lib.ErrorResult) {
	result, error := services.GetProducts()
	return result, error
}
