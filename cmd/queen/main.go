package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	message := `TYPE:RIDDLE
QUESTION:What has hands but can't clap?
END
`

	conn.Write([]byte(message))
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buffer[:n]))
}
