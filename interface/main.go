package main

import "fmt"

type Veiculo interface {
	start() string
}

type Car struct {
	Name string
}

type Motocicleta struct {
	Name string
}

func (c Car) start() string {
	return "Carro" + c.Name + " está iniciado"
}

func (m Motocicleta) start() string {
	return "Motocicleta" + m.Name + " está iniciado"
}

func startVeiculo(v Veiculo) string {
	return v.start()
}

func main() {
	c := Car{"Gol"}
	fmt.Println(c.start())

	m := Motocicleta{"XPTO"}
	fmt.Println(m.start())

	fmt.Println(startVeiculo(c))
	fmt.Println(startVeiculo(m))
}
