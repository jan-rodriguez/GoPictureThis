package database

import (
	// "database/sql"
	"github.com/jinzhu/gorm"

	"../models"
)

// CreateChallengeTable function to create challenges table, if it doens't exist
func CreateChallengeTable(db *gorm.DB) {
	// Create the challenges table
	if (!db.HasTable(&models.Challenge{})) {
		db.CreateTable(&models.Challenge{})
	}

	if !db.HasTable(&models.CreateChallenge{}) {
		db.CreateTable(&models.CreateChallenge{})
	}

}

// GetChallengesForUser gets the challenges created for a user
func GetChallengesForUser(db *gorm.DB, userID string, active bool) ([]models.Challenge, error) {
	var (
		challenges []models.Challenge
		// challenge  models.Challenge
		err error
	)

	// rows, err := db.Query(
	// 	`SELECT id, title, X(location), Y(location), picture_url, is_active, icon
	//        FROM challenges
	//        JOIN user_challenges
	//        ON user_challenges.challenge_id = challenges.id
	//        WHERE user_challenges.challenged_id = ?
	//        AND challenges.is_active=?
	//        GROUP BY challenges.id`, userID, active)

	// if err != nil {
	// 	return challenges, err
	// }

	// for rows.Next() {
	// 	err = rows.Scan(
	// 		&challenge.ID,
	// 		&challenge.Title,
	// 		&challenge.Location.Latitude,
	// 		&challenge.Location.Longitude,
	// 		&challenge.PictureURL,
	// 		&challenge.IsActive,
	// 		&challenge.Icon)
	// 	if err == nil {
	// 		challenges = append(challenges, challenge)
	// 	} else {
	// 		return challenges, err
	// 	}
	// }

	return challenges, err
}

// GetChallengesCreatedByUser retrieve list of challenges created by a user
func GetChallengesCreatedByUser(db *gorm.DB, userID string, active bool) ([]models.Challenge, error) {
	var (
		// challenge  models.Challenge
		challenges []models.Challenge
		err        error
	)

	// rows, err := db.Query(`
	//        SELECT id, title, X(location), Y(location), picture_url, is_active, icon
	//        FROM challenges
	//        WHERE challenger_id = ?
	//        AND is_active = ?`, userID, active)

	// if err != nil {
	// 	return challenges, err
	// }

	// for rows.Next() {
	// 	err = rows.Scan(
	// 		&challenge.ID,
	// 		&challenge.Title,
	// 		&challenge.Location.Latitude,
	// 		&challenge.Location.Longitude,
	// 		&challenge.PictureURL,
	// 		&challenge.IsActive,
	// 		&challenge.Icon)
	// 	if err == nil {
	// 		challenges = append(challenges, challenge)
	// 	} else {
	// 		return challenges, err
	// 	}
	// }

	return challenges, err
}

// CreateChallenge creates a challenge
func CreateChallenge(db *gorm.DB, json models.CreateChallenge) (models.Challenge, error) {
	var challenge models.Challenge
	var err error

	// Create the challenge
	// res, prepareErr := db.Exec(`
	//        INSERT INTO `+ChallengesTableName+`
	//        (title, challenger_id, location, picture_url, icon)
	//        VALUES
	//        (?, ?, Point(?, ?), ?, ?)`,
	// 	json.Title,
	// 	json.ChallengerID,
	// 	json.Location.Latitude,
	// 	json.Location.Longitude,
	// 	json.PictureURL,
	// 	json.Icon)

	// if prepareErr != nil {
	// 	return challenge, prepareErr
	// }

	// lastID, _ := res.LastInsertId()

	// challenge = models.Challenge{
	// 	ID:       int(lastID),
	// 	Icon:     json.Icon,
	// 	IsActive: true,
	// 	Location: models.Location{
	// 		Latitude:  json.Location.Latitude,
	// 		Longitude: json.Location.Longitude,
	// 	},
	// 	PictureURL: json.PictureURL,
	// 	Title:      json.Title,
	// }

	// // Now create all the challenges -> challenged mappings
	// stmt, err := db.Prepare(`
	//        INSERT INTO ` + UserChallengeTableName + `
	//        (challenge_id, challenger_id, challenged_id)
	//        VALUES
	//        (?, ?, ?)`)

	// defer stmt.Close()

	// for _, challengedID := range json.ChallengedIDs {

	// 	if err != nil {
	// 		return challenge, err
	// 	}

	// 	_, err = stmt.Exec(
	// 		lastID,
	// 		json.ChallengerID,
	// 		challengedID)

	// 	// TODO: Better error handling and rollback
	// 	if err != nil {
	// 		return challenge, err
	// 	}
	// }

	return challenge, err
}
