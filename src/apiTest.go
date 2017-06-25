package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func errorEval(err error, where string) bool {
	getCurrentTime := time.Now()
	if err != nil {
		fmt.Println(where+" ==> "+getCurrentTime.Format(time.RFC850)+":", err)
		return false
	}
	return true
}

func handlerIPTest(w http.ResponseWriter, r *http.Request) {
	err1 := db.Ping()
	if errorEval(err1, "Sql Ping") {
		data, err := getJSON("select * from sample limit ?,?", db)
		if err == nil {
			fmt.Printf(data)
			fmt.Fprintln(w, data)
		}
	}
}

func main() {

	db, err = openConnection("mysql", "root:L0k0t311@tcp(192.168.1.206:3306)/test")
	errorEval(err, "Sql Ping")
	err1 := db.Ping()
	errorEval(err1, "Sql Ping")

	http.HandleFunc("/test", handlerIPTest)

	http.ListenAndServe("localhost:4200", nil)
}
