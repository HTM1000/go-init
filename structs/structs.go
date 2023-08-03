package main

import "fmt"

type Car struct {
	name  string
	year  int
	color string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s\n", c.name, c.year, c.color)
}

func main() {
	car1 := Car{"Corolla", 2022, "white"}
	car2 := Car{"BMW", 2019, "blue"}

	fmt.Println(car1.name)
	fmt.Println(car2.year)

	fmt.Println(car1.info())

}
