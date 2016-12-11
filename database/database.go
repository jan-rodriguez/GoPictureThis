package database

import (
    "database/sql"
    "../models/"
)

const CHALLENGES_TABLE_NAME = "challenges"
const USERS_TABLE_NAME = "users"
const RESPONSES_TABLE_NAME = "responses"
const USER_CHALLENGE_TABLE_NAME = "user_challenges"

func CreateTables(db *sql.DB) error {
    // Create the challenges table
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS " + CHALLENGES_TABLE_NAME +
    `(
        id INT NOT NULL AUTO_INCREMENT,
        created TIMESTAMP NOT NULL DEFAULT now(),
        updated TIMESTAMP NOT NULL DEFAULT now() ON UPDATE now(),
        title varchar(100) NOT NULL,
        challenger_id INT NOT NULL,
        location POINT NOT NULL,
        picture_url varchar(100) NOT NULL,
        is_active BOOL NOT NULL DEFAULT 1,
        icon varchar(100) NOT NULL,
        PRIMARY KEY (id)
    )`)

    if (err != nil) {
        return err
    }

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS " + RESPONSES_TABLE_NAME +
        `(
            id INT NOT NULL AUTO_INCREMENT,
            created TIMESTAMP NOT NULL DEFAULT now(),
            updated TIMESTAMP NOT NULL DEFAULT now() ON UPDATE now(),
            challenge_id INT NOT NULL,
            user_id INT NOT NULL,
            status ENUM('open', 'accepted', 'declined', 'pending') NOT NULL,
            picture_url varchar(100) NOT NULL,
            PRIMARY KEY (id)
        )`)

    if (err != nil) {
        return err
    }

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS " + USERS_TABLE_NAME +
        `(
            id INT NOT NULL AUTO_INCREMENT,
            created TIMESTAMP NOT NULL DEFAULT now(),
            updated TIMESTAMP NOT NULL DEFAULT now() ON UPDATE now(),
            name varchar(100) NOT NULL,
            google_id varchar(100) NOT NULL,
            score INT NOT NULL DEFAULT 0,
            PRIMARY KEY (id)
        )`)

    if (err != nil) {
        return err
    }

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS " + USER_CHALLENGE_TABLE_NAME +
        `(
            challenge_id INT NOT NULL,
            challenger_id INT NOT NULL,
            challenged_id INT NOT NULL
        )`)

    return err
}

func GetChallengesForUser(db *sql.DB, user_id string, active bool) ([]models.Challenge, error)  {
    var (
        challenge models.Challenge
        challenges []models.Challenge
    )

    rows, err := db.Query("SELECT id, title, X(location), Y(location), picture_url, is_active, icon from " +
        CHALLENGES_TABLE_NAME + " WHERE challenger_id = ?", user_id)

    if (err != nil) {
        return challenges, err
    }

    for rows.Next() {
        err = rows.Scan(
            &challenge.Id,
            &challenge.Title,
            &challenge.Location.Latitude,
            &challenge.Location.Longitude,
            &challenge.Picture_Url,
            &challenge.Is_Active,
            &challenge.Icon)
        if (err == nil) {
            challenges = append(challenges, challenge)
        } else {
            return challenges, err
        }
    }

    return challenges, err
}

func CreateChallenge(db *sql.DB, json models.Create_Challenge) error {
    // Create the challenge
    res, prepareErr := db.Exec(`
        INSERT INTO ` + CHALLENGES_TABLE_NAME + `
        (title, challenger_id, location, picture_url, icon)
        VALUES
        (?, ?, Point(?, ?), ?, ?)`,
        json.Title,
        json.Challenger_Id,
        json.Location.Latitude,
        json.Location.Longitude,
        json.Picture_Url,
        json.Icon)

    if prepareErr != nil {
        return prepareErr
    }

    lastId, _ := res.LastInsertId()

    // Now create all the challenges -> challenged mappings
    stmt, err := db.Prepare(`
        INSERT INTO ` + USER_CHALLENGE_TABLE_NAME + `
        (challenge_id, challenger_id, challenged_id)
        VALUES
        (?, ?, ?)`)

    defer stmt.Close()

    for _, challenged_id := range json.Challenged_Ids {

        if err != nil {
            return err
        }

        _, err = stmt.Exec(
            lastId,
            json.Challenger_Id,
            challenged_id)

        // TODO: Better error handling and rollback
        if err != nil {
            return err
        }
    }

    return err
}
