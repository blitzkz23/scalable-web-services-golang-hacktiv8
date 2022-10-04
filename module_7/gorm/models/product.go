package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// * Saat struct ini dimigrate namanya otomatis menjadi plural
type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null;type:varchar(191)"`
	Brand string `gorm:"not null;type:varchar(191)"`
	// Ketika ada properti yang merupakan gabungan nama struct lain dengan ID (User + ID), maka gorm akan secara otomatis menganggapnya foreign key jadi tidak perlu diberikan tag seperti primayKey.
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// * Gorm hooks BeforeCreate
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product Before Create()")

	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}

	return
}
