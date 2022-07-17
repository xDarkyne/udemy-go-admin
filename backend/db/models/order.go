package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	OrderID    string      `json:"orderID" gorm:"type:text;primaryKey"`
	CreatedAt  time.Time   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	Total      float32     `json:"total" gorm:"-"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	Id         uint    `json:"id"`
	OrderID    string  `json:"order_id"`
	OrderTitle string  `json:"order_title"`
	Price      float32 `json:"price"`
	Quantity   uint    `json:"quantity"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	order.OrderID = uuid.NewString()
	return
}

func (order *Order) Count(orm *gorm.DB) int64 {
	var total int64
	orm.Model(&Order{}).Count(&total)

	return total
}

func (order *Order) Take(orm *gorm.DB, limit int, offset int) interface{} {
	var orders []Order
	orm.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	for i := range orders {
		var total float32 = 0
		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float32(orderItem.Quantity)
		}
		orders[i].Total = total
	}

	return orders
}
