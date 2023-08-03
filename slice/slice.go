package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)
	slice = append(slice, 1, 5, 7, 8) //adiciona um indice
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Println("Len:", len(slice), " Cap:", cap(slice))

	}

	sliceTest := slice
	slice[0] = 10
	slice = append(slice, 1, 5, 6, 9)
	fmt.Println(slice)
	fmt.Println(sliceTest)

	sliceString := [10]string{
		"Hello",
		"world",
		"much",
		"better",
		"better 2",
	}

	fmt.Println(sliceString)
	fmt.Println(sliceString[0])
	fmt.Println(sliceString[1:2])
	fmt.Println(sliceString[3:4])
	fmt.Println(sliceString[2:])
}
