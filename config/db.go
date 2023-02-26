package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() (err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", Config.DBHost, Config.DBUserName, Config.DBUserPassword, Config.DBName, Config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Logger = logger.Default.LogMode(logger.Info)

	//
	//log.Println("Running Migrations")
	//db.AutoMigrate(&model.User{})

	log.Println("🚀 Connected Successfully to the Database")
	DB = db
	return
}

//
//func Init(url string) Handler {
////	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
////
////	if err != nil {
////		log.Fatalln(err)
////	}
////
////	return Handler{db}
////}
