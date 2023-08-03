package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var result int
var m sync.Mutex

func main() {
	go runProcess("P1", 20)
	go runProcess("P2", 10)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Final result: ", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		m.Lock()
		result++
		m.Unlock()
		fmt.Println(name, " ->  ", i, " Partial Result: ", result)
	}
}
