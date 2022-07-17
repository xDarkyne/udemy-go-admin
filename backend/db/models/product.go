package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ProductID   string  `json:"productID" gorm:"type:text;primaryKey"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float32 `json:"price"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ProductID = uuid.NewString()
	return
}

func (product *Product) Count(orm *gorm.DB) int64 {
	var total int64
	orm.Model(&Product{}).Count(&total)

	return total
}

func (product *Product) Take(orm *gorm.DB, limit int, offset int) interface{} {
	var products []Product
	orm.Offset(offset).Limit(limit).Find(&products)
	return products
}
