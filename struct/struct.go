package main

import "log"

type User struct { // this is structure declaration
	Firstname string // define attributes and their types
	LastName  string
	Age       int
}

// Receivers
type Pen struct {
	CompanyName string
}

func (p *Pen) showCompanyName() string {
	return p.CompanyName
}

func main() {
	var usr User           // initialize a User struct
	usr.Firstname = "Lego" //assign attributes
	usr.LastName = "Jr"

	usr2 := User{ // shorthand to initialize struct and assign attributes
		Firstname: "Flash",
		LastName:  "Sr",
		Age:       21,
	}

	pen := Pen{
		CompanyName: "Pilot",
	}

	log.Println("User 1 :", usr)
	log.Println("User 2 :", usr2)

	log.Println(pen.CompanyName)
	log.Println(pen.showCompanyName())
}
