package database

import (
	"github.com/jinzhu/gorm"

	"../models/"
)

// CreateUsersTable : Creates user table if it doesn't exist
func CreateUsersTable(db *gorm.DB) {
	db.CreateTable(&models.User{})
}

// GetUserFromGoogleID : Gets user from google id
func GetUserFromGoogleID(db *gorm.DB, googleID string) (models.User, error) {
	var user models.User
	var err error

	// row := db.QueryRow(`
	//        SELECT id, name, score
	//        FROM users
	//        WHERE google_id = ?`, googleID)

	// err := row.Scan(
	// 	&user.ID,
	// 	&user.Name,
	// 	&user.Score)

	// user.GoogleID = googleID

	// return user, err
	return user, err
}

// CreateUser : Creates user
func CreateUser(db *gorm.DB, userJSON models.User) (*models.User, error) {
	var user models.User
	var err error

	// result, err := db.Exec(`
	//        INSERT INTO users
	//        (name, google_id)
	//        VALUES (?, ?)`, userJSON.Name, userJSON.GoogleID)

	// if err != nil {
	// 	return &user, err
	// }

	// lastID, err := result.LastInsertId()
	// user = models.User{
	// 	ID:       int(lastID),
	// 	Name:     userJSON.Name,
	// 	GoogleID: userJSON.GoogleID,
	// 	Score:    0,
	// }

	return &user, err
}
