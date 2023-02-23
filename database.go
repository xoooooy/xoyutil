package xoyutil

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnect(username string, password string, address string, port string, database string) *sql.DB {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+address+":"+port+")/"+database)
	errorPanic(err)

	return db
}

func dbClose(database *sql.DB) {
	database.Close()
}

func dbQuerry(database *sql.DB, sqlCode string) *sql.Rows {
	results, err := database.Query(sqlCode)
	errorPanic(err)

	return results
}

func sqlClean(sqlString string) string {
	sqlString = strings.ReplaceAll(sqlString, "'", "")
	sqlString = strings.ReplaceAll(sqlString, "\"", "")

	return sqlString
}
