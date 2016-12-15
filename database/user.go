package database

import (
    "database/sql"
    "../models/"
)

const USERS_TABLE_NAME = "users"

func CreateUsersTable(db *sql.DB) error {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS " + USERS_TABLE_NAME +
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

func GetUserFromGoogleId(db *sql.DB, google_id string) (models.User, error) {
    row := db.QueryRow(`
        SELECT id, name, score
        FROM users
        WHERE google_id = ?`, google_id)

    var user models.User

    err := row.Scan(
        &user.Id,
        &user.Name,
        &user.Score)

    return user, err
}

func CreateUser(db *sql.DB, user_json models.User) (*models.User, error) {
    var user models.User

    result, err := db.Exec(`
        INSERT INTO users
        (name, google_id)
        VALUES (?, ?)`, user_json.Name, user_json.Google_Id)

    if err != nil {
        return &user, err
    }

    last_id, err := result.LastInsertId()
    user = models.User {
        Id: int(last_id),
        Name: user_json.Name,
        Google_Id: user_json.Google_Id,
        Score: 0,
    }

    return &user, err
}
