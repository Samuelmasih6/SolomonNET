package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("King Solomon is listening...")

	fmt.Println("Waiting for Queen...")
	conn, err := listener.Accept()
	fmt.Println("Queen Connected!")

	fmt.Println("A queen has arrived!")

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("Message:", string(buffer[:n]))
}
