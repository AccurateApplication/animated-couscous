package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
JSON - standard notation for sending/recieving structured information. Can use others such as xml/asn.1 etc
standard pkgs: encoding/xml, encoding/json, encoding/asn1
Basic json types: numbers, bool, strings

*/

type Movie struct {
	Title       string
	Year        int  `json:"released"`        // field tag
	Color       bool `json:"color,omitempty"` // field tag. if false, no output produced
	Actors      []string
	notExported string
}

var movies = []Movie{
	{Title: "Movie1", Year: 1950, Color: false, Actors: []string{"forename1 surname1", "forename2 surname2"}, notExported: "hello"},
	{Title: "Movie500", Year: 2000, Color: true, Actors: []string{"actor actor actor", "forename2 surname2"}, notExported: "hello"},
}
var titles []struct{ Title string }

func main() {
	data, err := json.Marshal(movies)
	if err != nil {

		log.Fatalf("json marshal fail %s", err)
	}
	fmt.Printf("data:\n %s\n", data)

	if err := json.Unmarshal(data, &titles); err != nil { // here we unmarshal from data onto &titles. We only get Titles
		log.Fatalf("json unmarshal err: %s", err)

	}
	fmt.Println("title below")
	fmt.Println(titles)
}
