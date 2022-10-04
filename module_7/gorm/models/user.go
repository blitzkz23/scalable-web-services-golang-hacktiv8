package models

import (
	"time"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"not null;unique;type:varchar(191)"`
	// Asosiasi 1 user memiliki slice of product (banyak).  1 to many.
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
