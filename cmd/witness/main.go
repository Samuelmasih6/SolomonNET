package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	var message strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		// remove spaces/newlines
		line = strings.TrimSpace(line)

		if line == "END" {

			fmt.Println("Evidence Request:")
			fmt.Println(message.String())

			response := "TYPE:TESTIMONY\nANSWER:The merchant was 3.\nEND\n"

			_, err := conn.Write([]byte(response))
			if err != nil {
				return
			}

			message.Reset()
			continue
		}

		message.WriteString(line)
		message.WriteString("\n")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9093")
	if err != nil {
		panic(err)
	}

	fmt.Println("Witness is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}
