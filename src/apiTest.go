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

func handlerIPTests(response http.ResponseWriter, request *http.Request) {

	request.Header.Set("Content-type", "application/json")
	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	err1 := db.Ping()
	if errorEval(err1, "Sql Ping") {
		data, err := getJSON("select * from sample limit ?,?", 0, 100, db, 0)
		if err == nil {
			fmt.Printf(string(data))
			fmt.Fprintln(response, string(data))
		}
	}
}

func handlerIPTest(w http.ResponseWriter, r *http.Request) {
	err1 := db.Ping()
	if errorEval(err1, "Sql Ping") {
		data, err := getJSON("select id, name, salary  from sample WHERE ID=? LIMIT ?,?", 0, 100, db, 1)
		if err == nil {
			//fmt.Printf(data)
			fmt.Fprintln(w, data)
		}
	}
}

func main() {

	db, err = openConnection("mysql", "root:L0k0t311@tcp(192.168.1.206:3306)/test")
	errorEval(err, "Sql Ping")
	err1 := db.Ping()
	errorEval(err1, "Sql Ping")

	http.HandleFunc("/test", handlerIPTests)
	http.HandleFunc("/test?id=?", handlerIPTest)

	http.ListenAndServe("localhost:4200", nil)
}
