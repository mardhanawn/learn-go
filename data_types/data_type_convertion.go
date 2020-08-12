package main

import "fmt"

func main() {
	var x int32 = 10
	var y int64 = int64(x)

	fmt.Println(y)

	var z float64 = float64(y)

	fmt.Println(z)

	var b float64 = 2.5

	var a int32 = int32(b)

	fmt.Println(a)
}
