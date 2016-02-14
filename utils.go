package gofighter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadAPIKey(filename string)  (string, error) {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func MakeJSON (i interface{})  (string, error) {
	json_bytes, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "", err
	}
	json_str := string(json_bytes)
	return json_str, nil
}

func PrintJSON (i interface{}) {
	s, err := MakeJSON(i)
	if err != nil {
		return
	}
	fmt.Println(s)
	return
}
