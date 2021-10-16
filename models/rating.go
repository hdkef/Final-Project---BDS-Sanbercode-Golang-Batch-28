package models

import (
	"errors"

	"gorm.io/gorm"
)

type Rating struct {
	ID        uint `json:"id"`
	ArticleID uint `json:"article_id"`
	CreatorID uint
	Star      uint `json:"star"`
	gorm.Model
}

func (m *Rating) GetAll(db *gorm.DB, articleid int, lastid int, limit int) ([]Rating, error) {
	var ratings []Rating
	res := db.Where("article_id = ? AND id > ?", articleid, lastid).Find(&ratings).Limit(limit)
	return ratings, res.Error
}

func (m *Rating) Post(db *gorm.DB, articleid uint, usr *User) error {

	//validate must not empty
	if m.Star <= 0 {
		return errors.New("error ")
	}
	//validate rating
	if m.Star <= 0 || m.Star > 5 {
		return errors.New("rating star must be between 1 to 5")
	}
	//validate rating already rated article
	var rating Rating
	err := db.Where("creator_id = ? and article_id = ?", usr.ID, articleid).First(&rating).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("duplicate record")
	}

	//set creatorid and articleid
	m.CreatorID = usr.ID
	m.ArticleID = articleid
	return db.Create(m).Error
}

func (m *Rating) Put(db *gorm.DB, id int, usr *User) error {
	//validate rating
	if m.Star <= 0 || m.Star > 5 {
		return errors.New("rating star must be between 1 to 5")
	}

	var oldRating Rating
	err := db.Where("id = ?", id).First(&oldRating).Error
	if err != nil {
		return err
	}

	//validate if oldRating.CreatorID == usr.ID
	if oldRating.CreatorID != usr.ID {
		return errors.New("unauthorized")
	}
	return db.Model(&oldRating).Updates(*m).Error
}

func (m *Rating) Delete(db *gorm.DB, id int, usr *User) error {
	var rating Rating
	err := db.Where("id = ?", id).First(&rating).Error
	if err != nil {
		return err
	}

	//validate if rating.CreatorID == usr.ID
	if rating.CreatorID != usr.ID {
		return errors.New("unauthorized")
	}
	return db.Delete(&rating).Error
}
