package main

import "fmt"

func main() {
	var months = [12]string{
		"januari",
		"febriari",
		"maret",
		"aplil",
		"mey",
		"juni",
		"juli",
		"agusteus",
		"septemer",
		"oktober",
		"novemnber",
		"desember",
	}

	iniArray := [5]int{1, 2, 3, 4, 5} //array
	ininSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("ini array", iniArray)
	fmt.Println("ini slice", ininSlice)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "jhon"
	newSlice[1] = "doe"

	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy
	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)
	fmt.Println("ini yang diccopy ke newSlice", copySlice)

	// fmt.Println(copySlice)

	var monthFirst = months[5:7]

	fmt.Println(monthFirst)

	var monthSeccond = append(monthFirst, "jhon")
	fmt.Println(monthSeccond)

	monthFirst[1] = "bukan desember"
	fmt.Println(monthFirst)
}
