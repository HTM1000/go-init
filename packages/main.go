package main

import (
	"fmt"
	"main/packages/car"
)

func main() {
	car := car.Car{Name: "Gol", Color: "Black"}

	fmt.Println(car.Start())
}
