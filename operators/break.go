package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// fmt.Println(i) menampilkan dari 1 sampai dengan 100
		// if i%2 == 0 { // Menampilkan angka yang genap
		// 	continue // hasil dibawah perulangan pertama akan di skip
		// }

		if i == 50 {
			break
		}

		fmt.Println(i)
	}
}
