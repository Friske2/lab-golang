package main

import "fmt"

func main() {
	// useing var
	var name = "name2"
	var age int32 = 26
	const isCool = true
	firstName := "Parnupat"
	first, second := "one", "two"
	fmt.Println(name, age, first, second)
	fmt.Printf("%T\n", isCool) // return data type
	fmt.Println(firstName)
}
