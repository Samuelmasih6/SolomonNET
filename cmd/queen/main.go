package main

import (
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	message := `TYPE:RIDDLE
QUESTION:What has keys but can't open locks?
END`

	conn.Write([]byte(message))
}
