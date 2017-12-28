package dao

import (
	"encoding/json"
	"github.com/tsenart/nap"
)

//GetUser get all user info
func GetUser(db *nap.DB) string {
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	var arrResult map[string]interface{}
	var arrResults map[int]interface{}

	arrResult = make(map[string]interface{})
	arrResults = make(map[int]interface{})

	var i = 1
	for rows.Next() {
		i++
		var uid int
		var username string
		var department string
		var created int
		err = rows.Scan(&uid, &username, &department, &created)
		arrResult = make(map[string]interface{})
		arrResult["uid"] = uid
		arrResult["username"] = username
		arrResult["department"] = department
		arrResult["created"] = created

		arrResults[i] = arrResult
		checkErr(err)
	}
	mapRows, _ := json.Marshal(arrResults)
	return string(mapRows)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
