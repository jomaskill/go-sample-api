package dbConnection

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Start()  *gorm.DB{
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("Error from load .env file")
	}

	dsn := os.Getenv("dsn")
	dbType := os.Getenv("dbType")

	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed")
	}

	return db
}