package dao

import (
	"database/sql"
	"encoding/json"
)

func GetUser(db *sql.DB) string {
	rows, _ := fetchRows(db, "SELECT * FROM userinfo")
	mapRows, _ := json.Marshal(rows)
	return string(mapRows)
}