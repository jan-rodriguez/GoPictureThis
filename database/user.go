package database

import (
	"github.com/jinzhu/gorm"

	"../models/"
)

// CreateUsersTable : Creates user table if it doesn't exist
func CreateUsersTable(db *gorm.DB) {
	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
	}

	if !db.HasTable(&models.UserChallenge{}) {
		db.CreateTable(&models.UserChallenge{})
	}

}

// GetUserFromGoogleID : Gets user from google id
func GetUserFromGoogleID(db *gorm.DB, googleID string) (models.User, error) {
	var user models.User
	err := db.Where("google_id = ?", googleID).First(&user).Error
	return user, err
}

// CreateUser : Creates user
func CreateUser(db *gorm.DB, userJSON models.User) (*models.User, error) {
	existingUser, err := GetUserFromGoogleID(db, userJSON.GoogleID)
	// Already have existing user, just return that
	if existingUser.ID != 0 {
		return &existingUser, err
	}
	err = db.Create(&userJSON).Error
	return &userJSON, err
}
