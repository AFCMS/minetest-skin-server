package database

import (
	"log"
	"os"

	"minetest-skin-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var debug = os.Getenv("MSS_DEBUG") == "true"

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default,
	})
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//DB, err := gorm.Open(postgres.Open(fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%d", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), 5432)), &gorm.Config{
	//	Logger: logger.Default,
	//})

	if err != nil {
		log.Fatalln("Failed to connect database")
	}

	if res := DB.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		log.Fatalln(res.Error)
	}

	//defer DB.Close()

	err = DB.AutoMigrate(&models.Account{}, &models.Skin{}, &models.SkinLike{})
	if err != nil {
		log.Fatalln("Failed to migrate database")
	}
	//db.Logger = logger.Default.LogMode(logger.Info)
}
