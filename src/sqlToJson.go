package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func getJSON(sqlString string, page int, lines int, db *sql.DB, id int) ([]byte, error) {
	rows, err := db.Query(sqlString, 0, 10)
	fmt.Println(id)
	if id == 0 {
		rows, err = db.Query(sqlString, page, lines)
	} else {
		rows, err = db.Query(sqlString, id, page, lines)
	}

	if !errorEval(err, "db Query") {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
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
		return nil, err
	}
	//return string(jsonData), err
	return jsonData, err
}
