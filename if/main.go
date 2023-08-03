package main

import "fmt"

func main() {
	a := 10

	if a > 10 {
		fmt.Println("A > 10")
	} else if a > 5 {
		fmt.Println("A > 5")
	} else {
		fmt.Println("A < 10 e A < 5")
	}

	b := true
	c := true

	if x := "Cool"; b && c {
		fmt.Println(x)
	}

	if err := func() error {
		return nil
	}; err != nil {
	}
}
