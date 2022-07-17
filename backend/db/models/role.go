package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	RoleID      string       `json:"roleID" gorm:"type:text;primaryKey"`
	Name        string       `json:"name" gorm:"unique"`
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions"`
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.RoleID = uuid.NewString()
	return
}

func (role *Role) Count(orm *gorm.DB) int64 {
	var total int64
	orm.Model(&Role{}).Count(&total)

	return total
}

func (role *Role) Take(orm *gorm.DB, limit int, offset int) interface{} {
	var roles []Role
	orm.Preload("Permissions").Offset(offset).Limit(limit).Find(&roles)
	return roles
}
