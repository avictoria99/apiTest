package main

import "database/sql"

func openConnection(dbType string, connectionString string) (db *sql.DB, err error) {
	db, err = sql.Open(dbType, connectionString)
	return db, err
}
