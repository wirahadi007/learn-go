package main

import "fmt"

func main() {
	name := "wira"

	if name == "hadi" {
		fmt.Println(name)
	} else if name == "wira" {
		fmt.Println("hehehe salah")
	} else {
		fmt.Println("maaf tidak ada yang cocok dengan percabangan")
	}

	if length := len(name); length > 4 {
		fmt.Println("nama terlalu panjang")
	} else if length < 1 {
		fmt.Println("nama anda tidak benar")
	} else {
		fmt.Println("nama sudah benar")
	}
}
