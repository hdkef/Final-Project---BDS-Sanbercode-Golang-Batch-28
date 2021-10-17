package models

import (
	"errors"
	"net/url"

	"gorm.io/gorm"
)

type Draft struct {
	ID        uint   `json:"id"`
	CreatorID uint   `json:"creator_id"`
	Title     string `json:"title"`
	ImageURL  string `json:"image_url"`
	Body      string `json:"body"`
	gorm.Model
}

func (m *Draft) GetAll(db *gorm.DB, lastid int, limit int, usr *User) ([]Draft, error) {
	var drafts []Draft
	res := db.Where("id > ? AND creator_id = ?", lastid, usr.ID).Find(&drafts).Limit(limit)
	return drafts, res.Error
}

func (m *Draft) GetOne(db *gorm.DB, id uint, usr *User) (Draft, error) {
	var draft Draft
	if err := db.Where("id = ? AND creator_id = ?", id, usr.ID).First(&draft).Error; err != nil {
		return Draft{}, err
	}
	return draft, nil
}

func (m *Draft) Post(db *gorm.DB, usr *User) error {

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

func (m *Draft) Put(db *gorm.DB, id uint, usr *User) error {

	var oldDraft Draft
	err := db.Where("id = ?", id).First(&oldDraft).Error
	if err != nil {
		return err
	}

	//validate if usr.Role != "super-admin" then check if oldDraft.CreatorID == usr.ID
	if usr.Role != "super-admin" {
		if oldDraft.CreatorID != usr.ID {
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

	res := db.Model(&oldDraft).Updates(*m)
	return res.Error
}

func (m *Draft) Delete(db *gorm.DB, id uint, usr *User) error {
	var draft Draft
	if err := db.Where("id = ?", id).First(&draft).Error; err != nil {
		return err
	}
	//validate if usr.Role != "super-admin" then check if draft.CreatorID == usr.ID
	if usr.Role != "super-admin" {
		if draft.CreatorID != usr.ID {
			return errors.New("unauthorized")
		}
	}
	return db.Delete(&draft).Error
}
