package database

import (
	"strconv"

	"github.com/jinzhu/gorm"

	"../models/"
)

// ResponsesTableName responses table name
const ResponsesTableName = "responses"

// CreateResponsesTable exactly what it says
func CreateResponsesTable(db *gorm.DB) {
	if (!db.HasTable(&models.Response{})) {
		db.CreateTable(&models.Response{})
	}
}

func CreateResponse(db *gorm.DB, response models.Response) (models.Response, error) {
	response.Status = models.Open
	err := db.Create(&response).Error
	return response, err
}

// AcceptResponse accepts response
func AcceptResponse(db *gorm.DB, responseID string) (models.Response, error) {
	return updateResponseStatus(db, responseID, models.Accepted)
}

// DeclineResponse decline response
func DeclineResponse(db *gorm.DB, responseID string) (models.Response, error) {
	return updateResponseStatus(db, responseID, models.Declined)
}

func updateResponseStatus(db *gorm.DB, responseID string, status models.ResponseStatus) (models.Response, error) {
	idInt, err := strconv.ParseInt(responseID, 10, 0)

	response := models.Response{
		ID: int(idInt),
	}

	if err != nil {
		return response, err
	}

	err = db.Model(&response).Update("status", status).Error

	return response, err
}
