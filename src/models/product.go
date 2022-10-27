package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `gorm:"column:name; type:varchar(255); not null" json:"product_name"`
	Price int    `gorm:"column:price; not null" json:"product_price, omitempty"`
}
