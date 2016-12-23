package database

import (
	"database/sql"
)

// CreateTables create all the necessary tablaes for the game
func CreateTables(db *sql.DB) error {

	err := CreateChallengeTable(db)

	if err != nil {
		return err
	}

	err = CreateUsersTable(db)

	if err != nil {
		return err
	}

	err = CreateResponsesTable(db)

	if err != nil {
		return err
	}

	return err
}
