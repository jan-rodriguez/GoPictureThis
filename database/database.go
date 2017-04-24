package database

import (
	"github.com/jinzhu/gorm"
)

// CreateTables create all the necessary tablaes for the game
func CreateTables(db *gorm.DB) {

	CreateChallengeTable(db)

	// if err != nil {
	// 	return err

	// }
	// err = CreateUsersTable(db)

	// if err != nil {
	// 	return err
	// }

	// err = CreateResponsesTable(db)

	// if err != nil {
	// 	return err
	// }

	// return err
}
