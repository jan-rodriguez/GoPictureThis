package database

import (
	"github.com/jinzhu/gorm"

	"../models"
)

// CreateChallengeTable function to create challenges table, if it doens't exist
func CreateChallengeTable(db *gorm.DB) {
	// Create the challenges table
	if !db.HasTable(&models.Challenge{}) {
		db.CreateTable(&models.Challenge{})
	}
}

// GetChallengesForUser gets the challenges created for a user
func GetChallengesForUser(db *gorm.DB, userID string, active bool) ([]models.Challenge, error) {
	var challenges []models.Challenge

	err := db.Raw(
		`SELECT 
			id, 
			title, 
			latitude,
			longitude,
			picture_url,
			is_active,
			icon,
			challenges.challenger_id
	       FROM challenges
	       JOIN user_challenges
	       ON user_challenges.challenge_id = challenges.id
	       WHERE user_challenges.challenged_id = ?
	       AND challenges.is_active = ?
	       GROUP BY challenges.id`, userID, active).Scan(&challenges).Error

	return challenges, err
}

// GetChallengesCreatedByUser retrieve list of challenges created by a user
func GetChallengesCreatedByUser(db *gorm.DB, userID string, active bool) ([]models.Challenge, error) {
	var challenges []models.Challenge

	err := db.Where("challenger_id = ?", userID).Find(&challenges).Error

	return challenges, err
}

// CreateChallenge creates a challenge
func CreateChallenge(db *gorm.DB, createChallenge models.CreateChallenge) (models.Challenge, error) {
	var challenge models.Challenge

	challenge = models.Challenge{
		Icon:         createChallenge.Icon,
		IsActive:     true,
		Latitude:     createChallenge.Latitude,
		Longitude:    createChallenge.Longitude,
		PictureURL:   createChallenge.PictureURL,
		Title:        createChallenge.Title,
		IsGlobal: 	  createChallenge.IsGlobal,
		ChallengerID: createChallenge.ChallengerID,
	}

	err := db.Create(&challenge).Error

	for _, challengedID := range createChallenge.ChallengedIDs {

		if err != nil {
			return challenge, err
		}

		userChallenge := models.UserChallenge{
			ChallengeID:  challenge.ID,
			ChallengerID: createChallenge.ChallengerID,
			ChallengedID: challengedID,
		}

		err = db.Create(&userChallenge).Error

	}

	return challenge, err
}
