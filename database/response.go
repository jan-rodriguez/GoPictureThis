package database

import (
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

// AcceptResponse accepts response
func AcceptResponse(db *gorm.DB, responseID string) {
	updateResponseStatus(db, responseID, models.Accepted)
}

// DeclineResponse decline response
func DeclineResponse(db *gorm.DB, responseID string) {
	updateResponseStatus(db, responseID, models.Declined)
}

func updateResponseStatus(db *gorm.DB, responseID string, status models.ResponseStatus) {
	db.Exec(`
        UPDATE responses
        SET status = ?
        WHERE id = ?`, status.String(), responseID)
}
