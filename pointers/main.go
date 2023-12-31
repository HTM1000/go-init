package main

import "fmt"

func xpto(a *int) int {
	*a = 100
	return *a
}

func main() {
	b := 10
	fmt.Println(xpto(&b)) // 11
	fmt.Println(b)        // 10

	x := 10
	fmt.Println(&x) // 0xc00010e008

	y := &x
	fmt.Println(y)  // 0xc00010e008
	fmt.Println(*y) // 10

	*y = 20
	fmt.Println(x) // 20

	var z *int = &x
	fmt.Println(z)  // 0xc00010e008
	fmt.Println(*z) // 20
}
