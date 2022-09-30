package services

import (
	"github.com/awakelife93/gin-boilerplate/src/lib"
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/awakelife93/gin-boilerplate/src/models"
	"github.com/awakelife93/gin-boilerplate/src/repositories"
	"github.com/awakelife93/gin-boilerplate/src/utils"
)

func GetProducts() ([]models.Product, *structures.ErrorResult) {
	var products []models.Product
	repositories.ProductManager().Find(&products)
	return products, nil
}

func GetProduct(id int) (*models.Product, *structures.ErrorResult) {
	var product models.Product

	repositories.ProductManager().Where("id=?", id).First(&product)

	if utils.IsEmptyRecord(product.ID) {
		message := lib.NOT_FOUND_MESSAGE
		status := lib.NOT_FOUND_STATUS

		return nil, &structures.ErrorResult{
			Message: &message,
			Status:  &status,
		}
	}

	return &product, nil
}
