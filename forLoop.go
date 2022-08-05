package main

import "fmt"

func main() {
	conter := 1

	for conter <= 10 {
		fmt.Println("perulangan ke ", conter)
		conter++
	}

	//simple code for loop in go
	for i := 1; i <= 10; i++ {
		fmt.Println("perulangan ke ", i)
	}

	slice := []string{"kucing", "ayam", "anjing"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range
	names := []string{"kucing", "ayam", "anjing"}

	//bisa tanpa i jika tidak digunakan dengan nemmbahkan _ pada variable
	for i, value := range names {
		fmt.Println("index", i, "=", value)
	}

	person := make(map[string]string)

	person["name"] = "wirahadi"
	person["address"] = "denpasar"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
