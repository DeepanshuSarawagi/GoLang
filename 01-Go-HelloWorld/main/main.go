package main

import "fmt"

func main() {
	const name, age = "Kim", 22
	fmt.Println("Hello, Gophers 😊😁")
	fmt.Printf("%s is %d years old.\t And the type is %T and %T", name, age, name, age)
}
