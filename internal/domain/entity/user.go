package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name     string    `json:"name"`
	Username string    `json:"username" gorm:"unique"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `db:"password" json:"-"`
	Phone    string    `json:"phone" gorm:"type:varchar(13)"`
	Photo    string    `json:"photo"`
	Status   bool      `json:"status" gorm:"default:false"`
	CreateAt time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time `json:"update_at" gorm:"autoUpdateTime"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type ResponseLogin struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type UserUpdate struct {
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `db:"password" json:"-"`
	Phone    string `json:"phone" gorm:"type:varchar(13)"`
}
