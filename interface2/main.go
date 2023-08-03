package main

import "fmt"

type Name []interface {
}

func (n *Name) load() {
	*n = Name{
		"Henrique",
		"Neto",
		"Laura",
	}
}

func (n Name) printNames() {
	fmt.Println(n)
}

func Main() {
	var names Name
	names.load()
	names.printNames()
}
