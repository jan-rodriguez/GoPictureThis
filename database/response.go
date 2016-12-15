package database

import (
    "database/sql"
    "../models/"
)

const RESPONSES_TABLE_NAME = "responses"

func CreateResponsesTable(db *sql.DB) error {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS " + RESPONSES_TABLE_NAME +
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

func AcceptResponse(db *sql.DB, response_id string) error {
    return UpdateResponseStatus(db, response_id, models.Accepted)
}

func DeclineResponse(db *sql.DB, response_id string) error {
    return UpdateResponseStatus(db, response_id, models.Declined)
}

func UpdateResponseStatus(db *sql.DB, response_id string, status models.ResponseStatus) error {
    _, err := db.Exec(`
        UPDATE responses
        SET status = ?
        WHERE id = ?`, status.String(), response_id)
    return err
}
