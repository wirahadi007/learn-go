package main

import "fmt"

func main() {
	a := "ini typedata string dalam go"
	var (
		firstName = "wira"
		lastName  = "hadi"
	)
	var x int = 23
	var z string
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(x + len(a))

	z = "ini assignment to variable in golang"
	fmt.Println(z)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
