package controllers

import (
	"github.com/awakelife93/gin-boilerplate/src/lib"
	"github.com/awakelife93/gin-boilerplate/src/lib/structures"
	"github.com/awakelife93/gin-boilerplate/src/models"
	"github.com/awakelife93/gin-boilerplate/src/services"
	"github.com/awakelife93/gin-boilerplate/src/utils"
	"github.com/mitchellh/mapstructure"
)

func GetProducts(requestItem structures.RequestItem) (interface{}, *structures.ErrorResult) {
	result, error := services.GetProducts()
	return result, error
}

func GetProduct(requestItem structures.RequestItem) (interface{}, *structures.ErrorResult) {
	id := utils.GetRequestId(requestItem.Item)
	result, error := services.GetProduct(id)
	return result, error
}

func UpdateProduct(requestItem structures.RequestItem) (interface{}, *structures.ErrorResult) {
	id := utils.GetRequestId(requestItem.Item)
	var item *models.Product

	decodeError := mapstructure.Decode(requestItem.Item, &item)

	if decodeError != nil {
		message := decodeError.Error()
		status := lib.BAD_REQUEST_STATUS

		return nil, &structures.ErrorResult{
			Message: &message,
			Status:  &status,
		}
	}

	result, error := services.UpdateProduct(id, item)
	return result, error
}
