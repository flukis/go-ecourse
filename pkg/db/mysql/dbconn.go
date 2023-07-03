package mysql

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		panic("cannot read .env file")
	}

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DATABASE")
	uname := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")

	dsn := uname + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot open database connection")
	}

	return db
}
