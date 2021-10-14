package models

import (
	"errors"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint   `json:"id"`
	ArticleID uint   `json:"article_id"`
	CreatorID uint   `json:"creator_id"`
	Body      string `json:"body"`
	gorm.Model
}

func (m *Comment) GetAll(db *gorm.DB, articleid int, lastid int, limit int) ([]Comment, error) {
	var comments []Comment
	res := db.Where("article_id = ? AND id > ?", articleid, lastid).Find(&comments).Limit(limit)
	return comments, res.Error
}

func (m *Comment) Post(db *gorm.DB, articleID uint, usr *User) error {
	//validate must not empty
	err := MustNotEmpty(m.Body)
	if err != nil {
		return err
	}
	//set articleid and creatorid
	m.ArticleID = articleID
	m.CreatorID = usr.ID
	return db.Create(m).Error
}

func (m *Comment) Put(db *gorm.DB, id int, usr *User) error {
	//validate must not empty
	err := MustNotEmpty(m.Body)
	if err != nil {
		return err
	}

	var oldComment Comment
	err = db.Where("id = ?", id).First(&oldComment).Error
	if err != nil {
		return err
	}

	//validate if usr.Role != "super-admin" then check if oldComment.CreatorID == usr.ID
	if usr.Role != "super-admin" {
		if oldComment.CreatorID != usr.ID {
			return errors.New("unauthorized")
		}
	}
	return db.Model(&oldComment).Updates(*m).Error
}

func (m *Comment) Delete(db *gorm.DB, id int, usr *User) error {
	var comment Comment
	err := db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return err
	}

	//validate if usr.Role != "super-admin" then check if comment.CreatorID == usr.ID
	if usr.Role != "super-admin" {
		if comment.CreatorID != usr.ID {
			return errors.New("unauthorized")
		}
	}
	return db.Delete(&comment).Error
}
