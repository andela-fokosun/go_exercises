package main

import "fmt"

func main() {
	// var greet bool
	// var nogreet bool
	// var greater bool
	// var less bool

	var greet, nogreet, greater, less bool

	greet = true
	nogreet = false
	greater = 5 > 2
	less = 5 < 50

	fmt.Println(greet && nogreet)
	fmt.Println(greet || nogreet)
	fmt.Println(greater)
	fmt.Println(less)
}
