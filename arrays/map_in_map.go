package main

import "fmt"

func main() {
	// student := map[string]string{
	// 	"1": "Jerry",
	// 	"2": "Julia",
	// 	"3": "Joya",
	// 	"4": "Joyse",
	// 	"5": "Selly",
	// }

	student := map[string]map[string]string{
		"1": map[string]string{
			"name":    "Jerry",
			"address": "Banyumas",
			"grade":   "RPL 3",
		},
		"2": map[string]string{
			"name":    "Najib",
			"address": "Cilacap",
			"grade":   "RPL 5",
		},
		"3": map[string]string{
			"name":    "Abdul",
			"address": "Denpasar",
			"grade":   "RPL 9",
		},
		"4": map[string]string{
			"name":    "Abid",
			"address": "Pemalang",
			"grade":   "RPL 1",
		},
	}

	fmt.Println(student)
	fmt.Println(student["2"])
	fmt.Println(student["1"]["name"])
	fmt.Println(student["3"]["address"])
}
