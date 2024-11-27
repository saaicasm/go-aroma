package main

import (
	"log"
	"sort"
)

type User struct { // created a user type struct
	FisrtName string
	Lastname  string
}

func main() {

	//Map

	myMap := make(map[string]User) //create a map
	//map will lookup using a string and store a User type

	usr := User{ //created a usr of type User
		FisrtName: "Lego",
		Lastname:  "Jr",
	}

	myMap["me"] = usr //assigning string "me" to point to the usr struct

	log.Println(myMap)                 //the map will be map[me:{Lego Jr}]
	log.Println(myMap["me"])           //{Lego Jr}
	log.Println(myMap["me"].FisrtName) //Lego

	//Slices

	names := []string{"Lego", "Bobby", "Mo", "Lui"}

	log.Println(names)

	scores := []int{1, 3, 5, 7, 2, 4, 7}

	sort.Ints(scores)
	log.Println(scores)
}
