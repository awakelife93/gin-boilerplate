package services

import (
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/awakelife93/gin-boilerplate/src/models"
	"github.com/awakelife93/gin-boilerplate/src/repositories"
)

func GetProducts() (interface{}, *structures.ErrorResult) {
	var products models.Product
	repositories.ProductManager().Find(&products)
	return products, nil
}
