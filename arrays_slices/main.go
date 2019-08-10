package main

import "fmt"

func showFruit(value [2]string) {
	fmt.Println(value)
	fmt.Println(value[0])
}
func main() {
	//Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"
	showFruit(fruitArr)

	//Decalre and assign
	fruitArr2 := [2]string{"Apple", "Orange"}
	showFruit(fruitArr2)

	fruitArr3 := []string{"banana", "apple", "orange"}
	fmt.Println(fruitArr3)
	fmt.Println(len(fruitArr3)) // length of array
	fmt.Println(fruitArr3[1:2]) // show apple
}
