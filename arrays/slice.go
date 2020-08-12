package main

import "fmt"

func main() {
	names := [7]string{
		"Muhammad",
		"Ardhana",
		"Wahyu",
		"Nugraha",
		"Learn",
		"Go",
		"Language",
	}
	fmt.Println(names)

	fmt.Println(names[1:5])

	fmt.Println(names[:3])

	fmt.Println(names[2:])
}
