package models

import "gorm.io/gorm"

type Permission struct {
	Id   uint   `json:"id"`
	Name string `json:"name" gorm:"unique"`
}

func (permission *Permission) Count(orm *gorm.DB) int64 {
	var total int64
	orm.Model(&Permission{}).Count(&total)

	return total
}

func (permission *Permission) Take(orm *gorm.DB, limit int, offset int) interface{} {
	var permissions []Permission
	orm.Preload("Permissions").Offset(offset).Limit(limit).Find(&permissions)
	return permissions
}
