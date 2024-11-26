package pointer

import "log"

func main() {
	var myString string

	myString = "Red"

	log.Println("My string color is :", myString)

	log.Println(&myString)

}

// func changeUsingPointer(s x)
