package config

import (
	"bloggo/models"

	"gorm.io/gorm"
)

func initSuperAdmin(db *gorm.DB) error {
	var userCtx models.User = models.User{
		Role: "super-admin",
	}
	var superAdminUser models.User = models.User{
		Username:  "super-admin",
		Password:  "super-admin",
		AvatarURL: "https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg",
		Role:      "super-admin",
		Bio:       "i am super-admin",
	}
	return superAdminUser.Post(db, &userCtx)
}
