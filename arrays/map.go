package main

import "fmt"

func main() {
	student := make(map[string]string)

	student["324"] = "Ardhana"
	student["325"] = "Daffa"
	student["326"] = "Aziz"
	student["327"] = "Yulza"

	fmt.Println(student)
	fmt.Println(student["324"])
	fmt.Println(student["326"])
	fmt.Println(student["327"])
	fmt.Println(student["330"]) // tidak ada dalam data
}
