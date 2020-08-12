package main

import "fmt"

func main() {
	a := 9
	b := 2

	// result := a == b // a sama dengan b
	// result := a != b // a tidak sama dengan b
	// result := a < b // a kurang dari b
	// result := a > b // a lebih dari b
	// result := a <= b // a kurang dari atau sama dengan b
	result := a >= b // a lebih dari atau sama dengan b

	fmt.Println(result)
}
