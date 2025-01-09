package main

import (
	"fmt"
	"os"
)

func main() {

	a := os.Args[1]

	fmt.Printf("Hello, %s", a)
}
