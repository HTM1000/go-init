package main

import "fmt"

func main() {
	channel := make(chan int)
	channel <- 10 // deadlock -> tem que rodar dentro de uma goroutine
	go func() {
		channel <- 10
	}()
	fmt.Println(<-channel)
}
