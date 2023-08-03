package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"Model"` // troca o nome do elemento no json
	Year  int    `json:"-"`     // nao trazer o item no json
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
	Name   string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s\n", c.Name, c.Year, c.Color)
}

func main() {
	car := Car{"Gol", 2017, "Yellow"}

	result, _ := json.Marshal(car)
	fmt.Println(result)
	fmt.Println(string(result))

}
