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

func (m *User) GetOne(db *gorm.DB, userParam *User, id uint) (User, error) {
	//auth super-admin or usr.ID == id
	if userParam.Role != "super-admin" && userParam.ID != id {
		return User{}, errors.New("unauthorized")
	}
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (m *User) GetOnePublic(db *gorm.DB, id uint) (User, error) {

	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return User{}, err
	}
	//remove password
	var newUser User = User{
		ID:        user.ID,
		Username:  user.Username,
		Bio:       user.Bio,
		Role:      user.Role,
		AvatarURL: user.AvatarURL,
	}
	return newUser, nil
}

func (m *User) Post(db *gorm.DB, userParam *User) error {
	//validate must not empty
	err := MustNotEmpty(m.Username, m.Password, m.Bio, m.AvatarURL)
	if err != nil {
		return err
	}

	//validate only super-admin can post user
	if userParam.Role != "super-admin" {
		return errors.New("not super-admin")
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

func (m *User) GetAll(db *gorm.DB, usr *User, lastid int, limit int) ([]User, error) {
	//only super-admin can see all usr
	if usr.Role != "super-admin" {
		return nil, errors.New("not super-admin")
	}

	var users []User
	res := db.Where("id > ?", lastid).Find(&users).Limit(limit)
	return users, res.Error
}

func (m *User) Put(db *gorm.DB, usr *User, id uint) error {
	//only super-admin or usr.ID from token == id from path can edit
	if usr.Role != "super-admin" && usr.ID != id {
		return errors.New("unauthorized")
	}

	//cannot change role, password or username here
	var newUser User = User{
		Bio:       m.Bio,
		AvatarURL: m.AvatarURL,
	}

	//only super-admin can edit role
	if usr.Role == "super-admin" {
		newUser.Role = m.Role
	}

	var oldUser User
	err := db.Where("id = ?", id).First(&oldUser).Error
	if err != nil {
		return err
	}
	return db.Model(&oldUser).Updates(newUser).Error
}

func (m *User) Delete(db *gorm.DB, usr *User, id uint) error {
	//only super-admin or usr.ID from token == id from path can edit
	if usr.Role != "super-admin" && usr.ID != id {
		return errors.New("unauthorized")
	}
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	return db.Delete(&user).Error
}
