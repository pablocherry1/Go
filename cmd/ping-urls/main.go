package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(1 * time.Second) //Wait for 2 seconds
	fmt.Println("Hello from the goroutine!")
}

func main () {
	go sayHello()
	fmt.Println("Hello from the main function!")
	time.Sleep(2 * time.Second)
}