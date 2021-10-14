package config

import (
	"bloggo/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	_ = godotenv.Load()
}

func DBConn() (*gorm.DB, error) {

	var username string = os.Getenv("DBUSER")
	var password string = os.Getenv("DBPWD")
	var host string = os.Getenv("DBHOST")
	var database string = os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Article{}, &models.Comment{}, &models.Draft{}, &models.Inbox{}, &models.Media{}, &models.Rating{}, &models.User{})

	return db, nil
}
