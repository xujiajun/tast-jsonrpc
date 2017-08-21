// Package common implements json filename load and parse
package common

import (
	"encoding/json"
	io "io/ioutil"
)

// JSONStruct definition
type JSONStruct struct {
}

// NewJSONStruct be created new
func NewJSONStruct() *JSONStruct {
	return &JSONStruct{}
}

// Load JSON filename and parse
func (*JSONStruct) Load(filename string, v interface{}) {
	data, err := io.ReadFile(filename)
	if err != nil {
		return
	}
	jsonBytes := []byte(data)

	err = json.Unmarshal(jsonBytes, v)
	if err != nil {
		return
	}
}
