package database

import (
    "database/sql"

    "../models/"
)

// UsersTableName for users table
const UsersTableName = "users"

// CreateUsersTable : Creates user table if it doesn't exist
func CreateUsersTable(db *sql.DB) error {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS " + UsersTableName +
        `(
        id INT NOT NULL AUTO_INCREMENT,
        created TIMESTAMP NOT NULL DEFAULT now(),
        updated TIMESTAMP NOT NULL DEFAULT now() ON UPDATE now(),
        name varchar(100) NOT NULL,
        google_id varchar(100) NOT NULL,
        score INT NOT NULL DEFAULT 0,
        PRIMARY KEY (id),
        UNIQUE (google_id)
    )`)

    return err
}

// GetUserFromGoogleID : Gets user from google id
func GetUserFromGoogleID(db *sql.DB, googleID string) (models.User, error) {
    row := db.QueryRow(`
        SELECT id, name, score
        FROM users
        WHERE google_id = ?`, googleID)

    var user models.User

    err := row.Scan(
        &user.ID,
        &user.Name,
        &user.Score)

    user.GoogleID = googleID

    return user, err
}

// CreateUser : Creates user
func CreateUser(db *sql.DB, userJSON models.User) (*models.User, error) {
    var user models.User

    result, err := db.Exec(`
        INSERT INTO users
        (name, google_id)
        VALUES (?, ?)`, userJSON.Name, userJSON.GoogleID)

    if err != nil {
        return &user, err
    }

    lastID, err := result.LastInsertId()
    user = models.User{
        ID:       int(lastID),
        Name:     userJSON.Name,
        GoogleID: userJSON.GoogleID,
        Score:    0,
    }

    return &user, err
}
