package main

import "fmt"

func main() {
	nama := "wira"

	switch nama {
	case "wira":
		fmt.Println("hello", nama)
	case "hadi":
		fmt.Println("ini bukan wira")
	case "putra":
		fmt.Println("apalagi ini")

	default:
		fmt.Println("anda penipu")
	}

	//sort statement
	switch length := len(nama); length > 4 {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("hmmmm entahlah")
	}

	panjang := len(nama)

	switch {
	case panjang > 10:
		fmt.Println("nama anda terlalu panjang")
	case panjang < 2:
		fmt.Println("nama anda terlalu pendek")
	default:
		fmt.Println("nama sudah pas")
	}
}
