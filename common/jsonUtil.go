package common

import (
	"encoding/json"
	io "io/ioutil"
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
