package main

import "log"

func main() {
	var myString string // variable declaration

	myString = "Red" // assign a value

	log.Println("My string color is :", myString) 

	log.Println("Location of my string: ", &myString) //this prints hex of the mem add of variable

	changeUsingPointer(&myString) // call the func and pass the pointer to myVar
 
	log.Println("My string color after func call is :", myString) // this will print the updated value to the pointer
}

func changeUsingPointer(s *string) { //takes a parameter of type (pointer to string) s
	log.Println("The location of the string passed to func :", s) // this is the pointer to arg passed to the func 
	newValue := "Blue" // shorthand declaration of a neew variable
	*s = newValue //pointer to the 
}
