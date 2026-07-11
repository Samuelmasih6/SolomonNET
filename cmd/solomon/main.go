package main

import (
	"fmt"
	"net"
	"strings"
)

type Challenge struct {
	Type     string
	Question string
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	//fmt.Println("Received:", string(buffer[:n]))
	message := string(buffer[:n])

	challenge := parseChallenge(message)

	fmt.Println("Type:", challenge.Type)
	fmt.Println("Question:", challenge.Question)
}

func parseChallenge(message string) Challenge {
	var msgType string
	var question string

	lines := strings.Split(message, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "TYPE:") {
			msgType = strings.TrimPrefix(line, "TYPE:")
		}

		if strings.HasPrefix(line, "QUESTION:") {
			question = strings.TrimPrefix(line, "QUESTION:")
		}
	}

	return Challenge{
		Type:     msgType,
		Question: question,
	}
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("King Solomon is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}
