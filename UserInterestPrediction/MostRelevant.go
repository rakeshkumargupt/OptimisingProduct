package main

import (
        "fmt"
	"encoding/json"
)

func main(){
        
        fmt.Println("Starting ....")

	type Message struct {
	    Name string
	    Body string
	    Time int64
	}
	
	var mEncoding Message
	mEncoding = Message{"Alice", "Hello", 1294706395881547000}
	b := json.Marshal(mEncoding)
	fmt.Printf("%+v\n", b)

	// var mDecoding Message 
	// dec := json.Unmarshal(mDecoding, &mEncoding)
	// fmt.Printf("%+v\n", dec)
}
