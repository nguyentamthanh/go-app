package database

import (
	"github.com/tamthanh/go-app/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		panic("Failed connection to database")
	}
	log.Println("Connected Successfully")
	db.Logger = logger.Default.LogMode(logger.Info) //tao thoi gian moi
	log.Println("Running Migrations")

	db.AutoMigrate(&model.Product{})
	Database = DbInstance{
		Db: db,
	}
}
