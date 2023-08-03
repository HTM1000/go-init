package main

import "fmt"

func main() {
	a := "Henrique"
	switch a {
	case "Neto":
		fmt.Println("Olá Neto")
	case "Laura":
		fmt.Println("Olá Laura")
	default:
		fmt.Println("Não encontramos seu nome")
	}
}
