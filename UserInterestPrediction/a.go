package main

import (
        "fmt"
	"encoding/json"
)

	type Response1 struct {
		Page int
		Fruits []string
	}

	type Response2 struct {
		Page int `json:"page"`
		Fruits []string `json:"fruits"`	
	}
func main(){
        
        fmt.Println("Starting ....")

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	slcD :=   []string{"a","b","c"}
	slcB, _ :=  json.Marshal(slcD)
        fmt.Println(string(slcB))

	res1D := &Response1{
     	        Page:   1,
        	Fruits: []string{"apple", "peach", "pear"}}
        res1B, _ := json.Marshal(res1D)
        fmt.Println(string(res1B))

	res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])




















}

