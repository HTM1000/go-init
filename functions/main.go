package main

import "fmt"

func funcName(a int) int {
	return a * a
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func moreReturn(a string, b string) (string, string) {
	return a, b
}

func variadicFunctions(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}

	return res
}

func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {

	x := funcName(5)
	fmt.Println(x)

	fmt.Println(namedReturn("Henrique"))

	z, y := moreReturn("Henrique", "Neto")
	fmt.Println(z)
	fmt.Println(y)
	fmt.Println(variadicFunctions(1, 2, 5, 10))

	w := 0

	add := func() int {
		w += 2
		return w
	}

	fmt.Println(add())
	fmt.Println(add())
}
