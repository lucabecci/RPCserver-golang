package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	first := Item{"First", "a"}
	second := Item{"Second", "b"}
	third := Item{"Third", "c"}

	client.Call("API.AddItem", first, &reply)
	client.Call("API.AddItem", second, &reply)
	client.Call("API.AddItem", third, &reply)
	//db
	client.Call("API.GetDB", "", &db)

	fmt.Println("Database: ", db)
}
