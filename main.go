package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Witness preparing testimony...")
		ch <- "merchant"

		time.Sleep(2 * time.Second)

		fmt.Println("Witness delivered testimony")
	}()

	fmt.Println("Solomon waiting...")

	testimony := <-ch

	fmt.Println("Received:", testimony)
}
