package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// Welcome to channel your go guruji

// Topic go:embed
// go embed used to access static file within binary. 
// It was interoduce in go version 1.16
// var t []byte
//
//go:embed test.json
var d []byte

type Data struct {
	Name  string
	Phone string
}

func main() {
	fmt.Println(string(d))
	var dd Data

	json.Unmarshal(d, &dd)

	fmt.Println(dd.Phone)
}
