package main

import "log"

func main() {
	var myString string

	myString = "Red"

	log.Println("My string color is :", myString)

	log.Println("Location of my string: ", &myString)

	changeUsingPointer(&myString)

	log.Println("My string color after func call is :", myString)
}

func changeUsingPointer(s *string) {
	log.Println("The location of the string passed to func :", s)
	newValue := "Blue"
	*s = newValue
}
