package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	go func() {
		for {
			fmt.Print(<-channel)
		}
	}()

	time.Sleep(time.Second)
}
