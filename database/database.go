package database

import (
	"fmt"
	"log"
	"minetest-skin-server/models"
	"minetest-skin-server/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	// Allow error logging only if debug enabled
	var dblogger = logger.Default.LogMode(logger.Silent)
	if utils.ConfigDebugDatabase {
		dblogger = logger.Default
	}

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dblogger,
	})

	if err != nil {
		log.Fatalln("Failed to connect database")
	}

	err = DB.AutoMigrate(&models.Account{}, &models.Skin{}, &models.SkinLike{})
	if err != nil {
		log.Fatalln("Failed to migrate database")
	}
}
