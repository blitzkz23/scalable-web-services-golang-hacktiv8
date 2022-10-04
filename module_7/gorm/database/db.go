package database

import (
	"fmt"
	"log"
	"scalable-web-services-golang-hacktiv8/module_7/gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "naufalaldy23"
	dbPort   = 5432
	dbName   = "learning-gorm"
	db       *gorm.DB
	err      error
)

// * Menjalankan koneksi ke DB dan migrasi tabel.
func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, dbPort)

	// gorm.Open membangun koneksi kepada database.
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	// method debugging dichain dengan method automigrate
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

// * Mendapatkan referensi dari DB.
func GetDB() *gorm.DB {
	return db // Mendapatkan referensi *gorm.DB dari variabel db yang sudah direassign.
}
