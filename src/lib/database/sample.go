package database

import (
	"strconv"

	"github.com/awakelife93/gin-boilerplate/src/models"
)

func GenerateSampleData() {

	var products []models.Product
	GetDatabase().Find(&products)

	if len(products) < 1 {
		for i := 0; i < 100; i++ {
			products = append(products,
				models.Product{
					Name:  "Sample Product" + strconv.Itoa(i),
					Price: 100 + i,
				},
			)
		}

		// * bulk insert
		GetDatabase().CreateInBatches(products, 100)
	}
}
