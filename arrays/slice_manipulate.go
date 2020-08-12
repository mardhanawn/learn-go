package main

import "fmt"

func main() {
	// slice := []string{
	// 	"M",
	// 	"Muhammad",
	// 	"Ard",
	// 	"Ardha",
	// 	"Ardhana",
	// 	"Wahyu",
	// 	"N",
	// 	"Nugraha",
	// 	"MAWN",
	// 	"M Ardhana ",
	// }

	slice := make([]string, 3)
	slice[0] = "Wahyu"
	slice[1] = "Muhammad"
	slice[2] = "Ardha"

	fmt.Println(slice)
}
