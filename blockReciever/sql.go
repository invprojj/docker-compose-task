package main

import "database/sql"

const SqlCreateTable = "CREATE TABLE blocks(number VARCHAR PRIMARY KEY NOT NULL);"
const SqlInsertCock = "INSERT INTO blocks VALUES ($1);"

const (
	host     = "10.1.1.3"
	//host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func CreateDBConnection(connStr string)  {
	ps, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("failed to connect database")
	}
	db = *ps
}

func InsertBlockPostgre(value int64) {
	db.Exec(SqlInsertCock, value)
}

