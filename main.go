package main

import "log"

func main() {
	var myString string
	myString = "Lego"

	log.Println("This is my variable :", myString)

	log.Println(&myString)

	sh := &myString

	log.Println("This is pointer to the me add of my variable :", *sh)

}
