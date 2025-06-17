package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hz-evite/app/models"
)

func SetupDatabase() {
	var err error
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
	//	env.GetEnv("DB_HOST", "127.0.0.1"),
	//	env.GetEnv("DB_USER", ""),
	//	env.GetEnv("DB_PASSWORD", ""),
	//	env.GetEnv("DB_NAME", ""),
	//	env.GetEnv("DB_PORT", "5432"),
	//)
	fmt.Println(&DB)
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.User{}, &models.Words{})
	if err != nil {
		return
	}
}
