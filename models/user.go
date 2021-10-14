package models

import (
	"errors"
	"net/url"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
	gorm.Model
}

func (m *User) GetByUsername(db *gorm.DB) (User, error) {
	var usr User
	err := db.Where("username = ?", m.Username).First(&usr).Error
	if err != nil {
		return User{}, err
	}
	return usr, nil
}

func (m *User) Post(db *gorm.DB) error {
	//validate must not empty
	err := MustNotEmpty(m.Username, m.Password, m.Bio, m.AvatarURL)
	if err != nil {
		return err
	}

	//validate duplicate username
	var usr User
	err = db.Where("username = ?", m.Username).First(&usr).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("duplicate username")
	}

	//validate image_url
	_, err = url.ParseRequestURI(m.AvatarURL)
	if err != nil {
		return err
	}

	//hash password with bcrypt
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(m.Password), 5)
	if err != nil {
		return err
	}

	m.Password = string(hashedPass)

	return db.Create(m).Error
}
