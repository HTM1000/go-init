package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	u := uuid.NewV4()
	fmt.Printf("main, %s\n", u)
}
