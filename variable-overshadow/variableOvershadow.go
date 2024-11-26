package main

import "log"

var s = "from global"

func main() {
	var s2 string
	s2 = "from main"

	log.Println(shadowing("From Arguments"))
	log.Println(nonshadowing(s2))
}

func shadowing(s string) string {

	log.Println("The variable is from : ", s)
	return s
}
func nonshadowing(s2 string) string {

	log.Println("The variable is from : ", s)
	return s2
}
