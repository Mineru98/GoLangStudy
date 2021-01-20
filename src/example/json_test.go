package main

import (
	"encoding/json"
	"fmt"
)

type Res struct {
	Data	string	`json:"data"`
}

func json_test() {
	slice1 := []int{1,2,3}
    slice2 := append([]int{}, 4, 5)
	fmt.Println(slice1, slice2)
	
	x := make(map[string]int)
	x["key1"] = 10
	x["key2"] = 20
	fmt.Println(x)
	fmt.Println(x["key1"])
	fmt.Println(x["key2"])

	bolJson, _ := json.Marshal(true)
	fmt.Println(string(bolJson))

	res := &Res{
		Data: "test",
	}
	resJson, _ := json.Marshal(res)
	fmt.Println(string(resJson))
}