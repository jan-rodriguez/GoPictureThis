package database

import ("database/sql")


func CreateTables(db *sql.DB) error {

    err := CreateChallengeTable(db)

    if (err != nil) {
        return err
    }

    err = CreateUsersTable(db)

    if (err != nil) {
        return err
    }

    err = CreateResponsesTable(db)

    if (err != nil) {
        return err
    }


    return err
}
