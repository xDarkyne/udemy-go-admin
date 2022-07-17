package models

import "gorm.io/gorm"

type Entity interface {
	Count(orm *gorm.DB) int64
	Take(orm *gorm.DB, limit int, offset int) interface{}
}
