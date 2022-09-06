package controllers

import (
	"github.com/awakelife93/gin-boilerplate/src/lib"
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	"github.com/awakelife93/gin-boilerplate/src/services"
)

func GetNodes(item middleware.RequestItem) (interface{}, *lib.ErrorResult) {
	result, error := services.GetNodes()
	return result, error
}
