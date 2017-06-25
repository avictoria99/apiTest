package main

func openConnection(dbType string, connectionString as string) (db *sql.DB, err error) {
	db, err := sql.Open(dbType, connectionString)
	return db, err 
}
