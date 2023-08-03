package main

import "fmt"

type Car struct {
	name  string
	year  int
	color string
}

type SuperCar struct {
	Car
	canFly bool
	name   string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s\n", c.name, c.year, c.color)
}

func main() {
	car1 := Car{"Corolla", 2022, "white"}
	car2 := Car{"BMW", 2019, "blue"}
	scar := SuperCar{
		Car:    car1,
		canFly: true,
		name:   "Elantra",
	}

	fmt.Println(car1.name)
	fmt.Println(car2.year)

	fmt.Println(car1.info())
	fmt.Println(scar.info())

}
