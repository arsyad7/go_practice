package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Print() {
	fmt.Println("ID :", u.ID)
	fmt.Println("Email :", u.Email)
}
