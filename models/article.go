package models

import (
	"errors"
	"net/url"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint   `json:"id"`
	CreatorID uint   `json:"creator_id"`
	Title     string `json:"title"`
	ImageURL  string `json:"image_url"`
	Body      string `json:"body"`
}

func (m *Article) GetAll(db *gorm.DB) ([]Article, error) {
	var articles []Article
	res := db.Find(&articles)
	return articles, res.Error
}

func (m *Article) Post(db *gorm.DB, usr *User) error {

	m.CreatorID = usr.ID

	//validate must not empty
	err := MustNotEmpty(m.Body, m.Title, m.ImageURL)
	if err != nil {
		return err
	}

	//validate image_url
	_, err = url.ParseRequestURI(m.ImageURL)
	if err != nil {
		return err
	}

	return db.Create(m).Error
}

func (m *Article) Put(db *gorm.DB, id uint, usr *User) error {

	var oldArticle Article
	err := db.Where("id = ?", id).First(&oldArticle).Error
	if err != nil {
		return err
	}

	//validate if usr.Role != "super-admin" then check if oldArticle.CreatorID == usr.ID
	if usr.Role != "super-admin" {
		if oldArticle.CreatorID != usr.ID {
			return errors.New("unauthorized")
		}
	}

	//validate must not empty
	err = MustNotEmpty(m.Body, m.Title, m.ImageURL)
	if err != nil {
		return err
	}

	//validate image_url
	_, err = url.ParseRequestURI(m.ImageURL)
	if err != nil {
		return err
	}

	res := db.Model(&oldArticle).Updates(*m)
	return res.Error
}

func (m *Article) Delete(db *gorm.DB, id uint, usr *User) error {
	var article Article
	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		return err
	}
	//validate if usr.Role != "super-admin" then check if article.CreatorID == usr.ID
	if usr.Role != "super-admin" {
		if article.CreatorID != usr.ID {
			return errors.New("unauthorized")
		}
	}
	return db.Delete(&article).Error
}
