package main

import "fmt"

var name string
var age int

//-------------------------------------------------------------------
func displaythis(n string, a int) {
	var str_age string = fmt.Sprint(a)
	fmt.Println("Hello " + n)
	fmt.Println("You are " + str_age + " Years Old")
}

//-------------------------------------------------------------------
func main() {
	fmt.Print("Enter Your Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Your Age: ")
	fmt.Scanln(&age)

	displaythis(name, age)
}
