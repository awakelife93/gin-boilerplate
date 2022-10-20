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

func UpdateProduct(id int, item *models.Product) (*models.Product, *structures.ErrorResult) {
	product, getProductError := GetProduct(id)

	if getProductError != nil {
		return nil, getProductError
	}

	// todo: item 데이터 구조를 모델이 아닌 json 구조로...
	if item.Name != "" {
		product.Name = item.Name
	}

	if item.Price != 0 {
		product.Price = item.Price
	}

	sqlError := repositories.ProductManager().Updates(&product).Error

	if sqlError != nil {
		message := sqlError.Error()
		status := lib.INTERNAL_SERVER_ERROR_STATUS

		return nil, &structures.ErrorResult{
			Message: &message,
			Status:  &status,
		}
	}

	return product, nil
}
