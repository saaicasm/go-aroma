package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name  string
	Color string
}

func main() {

	dog := Dog{
		Name:  "Oscar",
		Breed: "Labrador",
	}

	gorilla := Gorilla{
		Name:  "Tungsten",
		Color: "Black",
	}

	PrintInfo(&dog)
	PrintInfo(&gorilla)

	log.Println(dog)            //{Oscar Labrador} it wont have Says()
	log.Println(gorilla.Says()) //this will have access to Says because of receiver
}

func PrintInfo(a Animal) {
	log.Println("The animal says ", a.Says(), "and has ", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string { //this will need a receiver to Dog
	return "woof woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Arghhhhh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
