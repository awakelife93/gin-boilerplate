package repositories

import (
	"fmt"

	"github.com/awakelife93/gin-boilerplate/src/lib/database"
	"github.com/awakelife93/gin-boilerplate/src/models"
	"gorm.io/gorm"
)

var productManager *gorm.DB = nil

func ProductManager() *gorm.DB {
	if productManager == nil {
		fmt.Println("productManager is nil")
	}

	return productManager
}

func setupProductManager() {
	productManager = database.GetDatabase().Model(&models.Product{})
}
