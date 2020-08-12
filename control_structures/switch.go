package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		// if i == 1 {
		// 	fmt.Println("Satu")
		// } else if i == 2 {
		// 	fmt.Println("Dua")
		// } else if i == 3 {
		// 	fmt.Println("Tiga")
		// } else if i == 4 {
		// 	fmt.Println("Empat")
		// } else if i == 5 {
		// 	fmt.Println("Lima")
		// } else {
		// 	fmt.Println("Lebih dari lima")
		// }

		switch i {
		case 1:
			fmt.Println("Satu")
		case 2:
			fmt.Println("Dua")
		case 3:
			fmt.Println("Tiga")
		case 4:
			fmt.Println("Empat")
		case 5:
			fmt.Println("Lima")
		default:
			fmt.Println("Lebih dari lima")
		}
	}
}
