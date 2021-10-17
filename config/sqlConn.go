package config

import (
	"bloggo/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	_ = godotenv.Load()
}

func DBConn() (*gorm.DB, error) {

	var username string = os.Getenv("DBUSER")
	var password string = os.Getenv("DBPWD")
	var host string = os.Getenv("DBHOST")
	var port string = os.Getenv("DBPORT")
	var database string = os.Getenv("DBNAME")

	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Article{}, &models.Comment{}, &models.Draft{}, &models.Inbox{}, &models.Media{}, &models.Rating{}, &models.User{})

	return db, nil
}
