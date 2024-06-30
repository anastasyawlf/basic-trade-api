package database

import (
	"basic-trade/models/entity"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	config := "root:@tcp(localhost:3306)/basic-trade?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(entity.Admin{}, entity.Product{}, entity.Variant{})

}

func GetDB() *gorm.DB {
	return db
}
