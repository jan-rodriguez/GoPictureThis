package mysql_helper

import ("github.com/go-sql-driver/mysql")

func GetMysqlCodeForError(err error) uint16 {
    if driverErr, ok := err.(*mysql.MySQLError); ok {
        // Now the error number is accessible directly
        return driverErr.Number
    } else {
        return 0
    }
}
