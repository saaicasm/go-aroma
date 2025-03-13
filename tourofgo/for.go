package main

import "fmt"

func main() {
	num := 1
	for num < 20 {
		fmt.Println(num)
		num += num
	}
}
