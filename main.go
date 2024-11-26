package main

import (
	"log"
)

type User struct {
	FirstName string
}

func (m *User) printName() string {
	log.Println("Check this :", m)
	return m.FirstName
}
func main() {
	var myUsr User
	myUsr.FirstName = "Lego"

	myUsr2 := User{}

	log.Println("My User 1 is :", myUsr.printName())
	log.Println("This should be default: ", myUsr2.printName())
	// log.Println("My User 2 is :", myUsr2.printName())
}
