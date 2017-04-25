package database

import (
	"github.com/jinzhu/gorm"
)

// CreateTables create all the necessary tablaes for the game
func CreateTables(db *gorm.DB) {

	CreateChallengeTable(db)
	CreateUsersTable(db)
	CreateResponsesTable(db)
}
