package main

import "fmt"

func main() {

	/*
		In this section, we are going to learn about variables in GoLang. A VARIABLE holds a value and variable type.
		Since, GoLang is a statically typed language, we need to define what type of value a variable is going to hold.

		// Refer to the example below.
	*/

	var a = 43
	fmt.Println(a)
	// Now as you can see below commented code will not work. It will complain "Type String cannot be represented by type int"
	// a = "Deepanshu"

	// However, this re-assignment of same type should work.
	a = 57

	fmt.Printf("The value of a is now %d", a)

	/*
		We can also use the short declaration operator here in GoLang to declare a variable and assign values to them.
		If one needs to assign multiple variables together, then it can be achieved using short declaration operator as seen below.
	*/

	b := 100 // Short declaration operator
	fmt.Println(b)

	c, d, e, f := 15, 95, 190, "Deepanshu" // Multiple assignment
	fmt.Println(c, d, e, f)

}
