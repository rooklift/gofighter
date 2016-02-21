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

func MakeJSON(i interface{})  (string, error) {
	json_bytes, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "", err
	}
	json_str := string(json_bytes)
	return json_str, nil
}

func PrintJSON(i interface{})  {
	s, err := MakeJSON(i)
	if err != nil {
		return
	}
	fmt.Println(s)
	return
}

func MoveFromOrder(order Order)  Movement {

	// Assuming the order is closed, this returns the total
	// change in cents and shares caused by it...

	var m Movement

	m.Shares = order.TotalFilled
	if order.Direction == "sell" {
		m.Shares *= -1
	}

	if order.Fills != nil {
		for _, fill := range order.Fills {
			if order.Direction == "sell" {
				m.Cents += fill.Qty * fill.Price
			} else {
				m.Cents -= fill.Qty * fill.Price
			}
		}
	}

	return m
}
