package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
func mutilValues(name string, age int) (string, int) {
	return name, age
}
func main() {
	fmt.Println(greeting("Kit"))
	fmt.Println(getSum(1, 1))
	fmt.Println(mutilValues("kit", 25))
}
