package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	
		[
			{
				"first_name":"Roberto",
				"last_name":"Firmino",
				"has_dog":false
			},
			{
				"first_name":"Kostas",
				"last_name":"Tsimikas",
				"has_dog":true
			}
		]

	`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("The error is : ", err)
	}

	log.Printf("Unmarshalling : %v", unmarshalled)

	// write json form a struct

	var mySlice []Person

	var m1 Person

	m1.FirstName = "Emre"
	m1.LastName = "Can"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person

	m2.FirstName = "Virgil"
	m2.LastName = "Dijk"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, " ", "  ")

	if err != nil {
		log.Println(err)
	}

	log.Println(string(newJson))

}
