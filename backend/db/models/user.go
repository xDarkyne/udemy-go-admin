package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID       string    `json:"userID" gorm:"type:text;primaryKey"`
	Username     string    `json:"username" gorm:"unique"`
	Email        string    `json:"email" gorm:"unique"`
	Password     string    `json:"-"`
	CreatedAt    time.Time `json:"-" gorm:"autoCreateTime"`
	LastLoggedIn time.Time `json:"-"`
	RoleID       string    `json:"roleID"`
	Role         Role      `json:"role" gorm:"foreign_key:RoleID" `
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.UserID = uuid.NewString()
	return
}

func (user *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) IsValidPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (user *User) Count(orm *gorm.DB) int64 {
	var total int64
	orm.Model(&User{}).Count(&total)

	return total
}

func (user *User) Take(orm *gorm.DB, limit int, offset int) interface{} {
	var users []User
	orm.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	return users
}
