package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Challenge struct {
	Type     string
	Question string
}

var evidenceDB = map[string]string{
	"Who stole the treasure?": "MERCHANT",
	"Who stole the crown?":    "GUARD",
	"Who burned the barn?":    "FARMER",
}

func findSuspect(question string) string {

	suspect, ok := evidenceDB[question]

	if !ok {
		return "UNKNOWN"
	}

	return suspect
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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	var message strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		line = strings.TrimSpace(line)

		if line == "END" {

			challenge := parseChallenge(message.String())

			suspect := findSuspect(
				challenge.Question,
			)

			response := fmt.Sprintf(
				"TYPE:TESTIMONY\nANSWER:%s\nEND\n",
				suspect,
			)

			conn.Write([]byte(response))

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
