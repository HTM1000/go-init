package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
	Name   string
}

func main() {
	var car Car
	data := []byte(`{"Name": "Gol", "Year": 2000, "Color": "white"}`)

	json.Unmarshal(data, &car)
	fmt.Println(car.Name, car.Year, car.Color)

}
