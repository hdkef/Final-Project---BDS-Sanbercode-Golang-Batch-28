package models

import (
	"errors"

	"gorm.io/gorm"
)

type Inbox struct {
	ID         uint   `json:"id"`
	CreatorID  uint   `json:"creator_id"`
	ReceiverID uint   `json:"receiver_id"`
	Body       string `json:"body"`
	gorm.Model
}

func (m *Inbox) GetAll(db *gorm.DB, receiverid int, lastid int, limit int) ([]Inbox, error) {
	var inboxs []Inbox
	res := db.Where("receiver_id = ? AND id > ?", receiverid, lastid).Find(&inboxs).Limit(limit)
	return inboxs, res.Error
}

func (m *Inbox) Post(db *gorm.DB, receiverID uint, usr *User) error {
	//cannot send msg to own inbox
	if receiverID == usr.ID {
		return errors.New("cannot send msg to own inbox")
	}

	//validate must not empty
	err := MustNotEmpty(m.Body)
	if err != nil {
		return err
	}
	//set receiverid and creatorid
	m.ReceiverID = receiverID
	m.CreatorID = usr.ID
	return db.Create(m).Error
}

func (m *Inbox) Put(db *gorm.DB, id int, usr *User) error {
	//validate must not empty
	err := MustNotEmpty(m.Body)
	if err != nil {
		return err
	}

	var oldInbox Inbox
	err = db.Where("id = ?", id).First(&oldInbox).Error
	if err != nil {
		return err
	}

	//validate if usr.Role != "super-admin" then check if oldInbox.CreatorID == usr.ID
	if usr.Role != "super-admin" {
		if oldInbox.CreatorID != usr.ID {
			return errors.New("unauthorized")
		}
	}
	return db.Model(&oldInbox).Updates(*m).Error
}

func (m *Inbox) Delete(db *gorm.DB, id int, usr *User) error {
	var inbox Inbox
	err := db.Where("id = ?", id).First(&inbox).Error
	if err != nil {
		return err
	}

	//validate if usr.Role != "super-admin" then check if inbox.CreatorID == usr.ID
	if usr.Role != "super-admin" {
		if inbox.CreatorID != usr.ID {
			return errors.New("unauthorized")
		}
	}
	return db.Delete(&inbox).Error
}
