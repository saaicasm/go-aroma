package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Lego")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{
		"Gupt",
		"Shatrugan",
		"Baba",
	}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range messages {
		fmt.Println(msg)
	}

}
