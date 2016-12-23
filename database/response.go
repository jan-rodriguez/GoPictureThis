package database

import (
	"database/sql"

	"../models/"
)

// ResponsesTableName responses table name
const ResponsesTableName = "responses"

// CreateResponsesTable exactly what it says
func CreateResponsesTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + ResponsesTableName +
		`(
        id INT NOT NULL AUTO_INCREMENT,
        created TIMESTAMP NOT NULL DEFAULT now(),
        updated TIMESTAMP NOT NULL DEFAULT now() ON UPDATE now(),
        challenge_id INT NOT NULL,
        user_id INT NOT NULL,
        status ENUM('open', 'accepted', 'declined') NOT NULL,
        picture_url varchar(100) NOT NULL,
        PRIMARY KEY (id)
    )`)

	return err
}

// AcceptResponse accepts response
func AcceptResponse(db *sql.DB, responseID string) error {
	return updateResponseStatus(db, responseID, models.Accepted)
}

// DeclineResponse decline response
func DeclineResponse(db *sql.DB, responseID string) error {
	return updateResponseStatus(db, responseID, models.Declined)
}

func updateResponseStatus(db *sql.DB, responseID string, status models.ResponseStatus) error {
	_, err := db.Exec(`
        UPDATE responses
        SET status = ?
        WHERE id = ?`, status.String(), responseID)
	return err
}
