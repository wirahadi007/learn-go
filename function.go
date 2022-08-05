package main

import "fmt"

//function lengkap dengan parameternya yang ada

func main() {
	// getLoop()

	sayHello("wirahadi")
	fmt.Println(getHello("wirahadi"))

	//ignoring beberapa value dengan _

	firstName, _, lastName := getFullName()

	// firstName, _, lastName := getFullName()

	fmt.Println(firstName)
	// fmt.Println(middleName)
	fmt.Println(lastName)
}

func getLoop() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		sayHello("wirahadi")
	}
}

func sayHello(name string) {
	fmt.Print("hello", " ", name)

}

//return value singel
func getHello(name string) string {
	return "Heloo" + name
}

//return multiple values
func getFullName() (string, string, string) {
	return "wira", "hadi", "putra"
}

//named return values

func getComplatedName() (firstName, middleName, lastName string) {
	firstName = "wira"
	middleName = "hadi"
	lastName = "putra"

	//tidak wajib
	// return firstName, middleName, lastName
	return

}
