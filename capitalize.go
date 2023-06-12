package main

import (
	"fmt"
	"strings"
)

func capitalizeThis(name string) {
	allcaps_name := strings.ToUpper(name)

	fmt.Println("Your name is " + allcaps_name)
}

func main() {
	var name string

	fmt.Print("Enter Your Name: ")
	fmt.Scanln(&name)

	capitalizeThis(name)
}
