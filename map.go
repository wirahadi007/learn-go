package main

import "fmt"

func main() {
	person := map[string]string{
		"firstName": "wirahadi",
		"lastName":  "putra",
		"address":   "denpasar",
	}

	fmt.Println(person)
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])
	fmt.Println(person["address"])

	//manipulasi map
	fmt.Println(len(person))

	person["address"] = "jembrana"
	person["firstName"] = "putra"
	person["lastname"] = "hadi"

	fmt.Println(person["address"])

	delete(person, "address")

	fmt.Println(len(person))

	book := make(map[string]string)
	book["name"] = "jhon"
	book["penerbit"] = "entahlan"

	fmt.Println(book)
	fmt.Println(book["name"])

	book["name"] = "doe"

	fmt.Println(book["name"])

}
