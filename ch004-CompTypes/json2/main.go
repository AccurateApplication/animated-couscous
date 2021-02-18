package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
	resp, err := http.Get("https://xkcd.com/121/info.0.json")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(resp.Body)
	// var testing []testing123
	// json.Unmarshal(resp.Body, &testing)
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
