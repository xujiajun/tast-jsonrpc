package common

import (
	io "io/ioutil"
	"encoding/json"
)

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (self *JsonStruct) Load(filename string, v interface{}) {
	data, err := io.ReadFile(filename)
	if err != nil {
		return
	}
	datajson := []byte(data)

	err = json.Unmarshal(datajson, v)
	if err != nil {
		return
	}
}
