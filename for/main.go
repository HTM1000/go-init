package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for {
		x++

		if x == 50 {
			continue
		}

		if x == 100 {
			break
		}
	}

}
