package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func errorEval(err error, where string) bool {
	getCurrentTime := time.Now()
	if err != nil {
		fmt.Println(getCurrentTime.Format(time.RFC850)+":", err)
		return false
	}
	return true
}

func getJSON(sqlString string, db *sql.DB) (string, error) {
	rows, err := db.Query(sqlString, 0, 10)
	if !errorEval(err, "db Query") {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(jsonData))
	return string(jsonData), err
}

func handlerIPTest(db *sql.DB) {
	err1 := db.Ping()
	if errorEval(err1, "Sql Ping") {
		//data, err := db.Query("select * from sample limit ?",100000)
		data, err := getJSON("select * from sample limit ?,?", db)
		if err == nil {
			fmt.Printf(data)
			//w.WriteHeader(http.StatusOK)
			//w.Write(data)
			//fmt.Fprintln(w, data)
		}
	}
}

func main() {

	db, err := storage.openConnection("mysql", "root:L0k0t311@tcp(192.168.1.206:3306)/test")

	http.HandleFunc("/test", handlerIPTest(db))

	http.ListenAndServe("localhost:4200", nil)

	/*	for data.Next() {
			var (
				id          string
				name, depto string
				salary      string
				enterDate   string
			)

			err := data.Scan(&id, &name, &depto, &salary, &enterDate)
			fmt.Println(id + " " + name + " " + depto + " " + salary)
			if err != nil {
				fmt.Println("Error connection")
			}
		}
	*/

}
