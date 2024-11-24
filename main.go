package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var saySmth string
	var i int = 1
	saySmth = "Goodbye"

	fmt.Println(saySmth)
	fmt.Println("is is i ", i)

	wht := sayNo()

	fmt.Println(wht)

}

func sayNo() string {
	return "Something"
}
