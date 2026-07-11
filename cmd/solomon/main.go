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
			challenge := parseChallenge(message.String())

			fmt.Println("Type:", challenge.Type)
			fmt.Println("Question:", challenge.Question)

			answer := solveRiddle(challenge.Question)
			response := fmt.Sprintf("TYPE:ANSWER\nANSWER:%s\nEND\n", answer)
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

func solveRiddle(question string) string {
	switch question {
	case "What has keys but can't open locks?":
		return "A piano"

	case "What has hands but can't clap?":
		return "A clock"

	default:
		return "I do not know"
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
