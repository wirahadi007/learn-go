package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noktpW noKTP = "23123123123123"
	fmt.Println(noktpW)

	var marriedSttatus Married = false

	fmt.Println(marriedSttatus)
}
