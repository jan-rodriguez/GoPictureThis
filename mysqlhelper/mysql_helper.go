package mysqlhelper

import ("github.com/go-sql-driver/mysql")

// GetMysqlCodeForError gets error number for mysql error, 0 if not a sql error
func GetMysqlCodeForError(err error) uint16 {
    if driverErr, ok := err.(*mysql.MySQLError); ok {
        // Now the error number is accessible directly
        return driverErr.Number
    }
    
    return 0
}
