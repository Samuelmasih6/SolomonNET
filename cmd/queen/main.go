package main

import (
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("Greetings King Solomon"))
}
