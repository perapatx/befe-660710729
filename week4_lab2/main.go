package main

import (
	"fmt"
)

func main() {

	// var name string = "perapat"
	var age int = 25
	email := "mok@gmail.com"
	gpa := 3.8
	firstname, lastname := "perapat", "klaysamneang"

	fmt.Printf("name: %s \nsurname: %s \nage: %d\nemail: %s\ngpa: %.2f\n", firstname, lastname, age, email, gpa)

}
