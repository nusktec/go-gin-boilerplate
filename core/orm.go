package core

import (
	"fmt"
	"goginapi/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Orm struct {
}

func (core Orm) Done() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	pass := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_DATABASE")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to find .env file")
	}

	// migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Book{},
	)
}
