package database

import (
	"log"

	"minetest-skin-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	DB, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		log.Fatalln("Failed to connect database")
	}

	if res := DB.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		log.Fatalln(res.Error)
	}

	//defer DB.Close()

	err = DB.AutoMigrate(&models.Account{}, &models.Skin{})
	if err != nil {
		log.Fatalln("Failed to migrate database")
	}
	//db.Logger = logger.Default.LogMode(logger.Info)
}
