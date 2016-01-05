package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page   int
	Fruits []string
}

func main() {
	//fmt.Println("HI there")
	res := Response{Page: 1, Fruits: []string{"apple", "peach"}}
	res_json, err := json.Marshal(&res)
	if err != nil {
		panic("json marshalling failed")
	}
	fmt.Println(string(res_json))
}
