package models

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Media struct {
	ID        uint   `json:"id"`
	CreatorID uint   `json:"creator_id"`
	Caption   string `json:"caption"`
	Alt       string `json:"alt"`
	Url       string `json:"url"`
	gorm.Model
}

func init() {
	godotenv.Load()
}

var MEDIAPTH string = os.Getenv("MEDIAPTH")

func (m *Media) GetAll(db *gorm.DB, usr *User, id uint, lastid int, limit int) ([]Media, error) {
	//validate creator_id = usr.id and user is not super-admin
	if usr.ID != id && usr.Role != "super-admin" {
		return nil, errors.New("unauthorized")
	}

	var media []Media
	res := db.Where("creator_id = ? AND id > ?", id, lastid).Find(&media).Limit(limit)
	return media, res.Error
}

func (m *Media) Post(db *gorm.DB, usr *User) error {
	//set creatorID
	m.CreatorID = usr.ID

	//validate empty
	err := MustNotEmpty(m.Alt, m.Caption)
	if err != nil {
		return err
	}

	return db.Create(m).Error
}

func (m *Media) Delete(db *gorm.DB, usr *User, id uint) error {
	var media Media
	err := db.Where("id = ?", id).First(&media).Error
	if err != nil {
		return err
	}

	//validate if media.CreatorID == usr.ID
	if media.CreatorID != usr.ID {
		return errors.New("unauthorized")
	}

	//delete from disk
	relPath := MEDIAPTH + media.Url
	err = os.Remove(relPath)
	if err != nil {
		return err
	}

	//delete from db
	return db.Delete(&media).Error
}
