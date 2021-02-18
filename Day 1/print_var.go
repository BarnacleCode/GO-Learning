package main

import "fmt"

func main() {
	var str string

	//assign values
	str = "Hello world"
	age := 20

	fmt.Println(str)
	fmt.Println(age)

	//multiple declaration
	var (
		name     string
		email    string
		children int
	)

	fmt.Printf("Your name :%s\n", name)
	fmt.Printf("Your Email:%s\n", email)
	fmt.Printf("How many children do you have?%d", children)

}
