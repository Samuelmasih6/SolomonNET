package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := `TYPE:RIDDLE
QUESTION:What has keys but can't open locks?
END
`

	_, err = conn.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)

	var response strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)

		if line == "END" {
			fmt.Println(response.String())
			break
		}

		response.WriteString(line)
		response.WriteString("\n")
	}
}
