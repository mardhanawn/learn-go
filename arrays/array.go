package main

import "fmt"

func main() {
	var names [6]string

	names[0] = "Muhammad"
	names[1] = "Ardhana"
	names[2] = "Wahyu"
	names[3] = "Nugraha"
	names[4] = "Learn"
	names[5] = "Go"

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println("index", index, "=", value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	// fmt.Println(names) // Mengambil semua data array
	// fmt.Println(names[4]) // Mengambil index data ke-x array

}
