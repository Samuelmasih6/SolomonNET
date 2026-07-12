package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	consolereader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(conn)

	for {
		fmt.Print("Challenge Type (RIDDLE/CASE):")
		challengetype, err := consolereader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Print("Question:")
		question, err := consolereader.ReadString('\n')
		if err != nil {
			return
		}

		challengetype = strings.TrimSpace(challengetype)
		question = strings.TrimSpace(question)

		if question == "exit" {
			return
		}

		message := fmt.Sprintf(
			"TYPE:%s\nQUESTION:%s\nEND\n",
			challengetype,
			question,
		)
		_, err = conn.Write([]byte(message))
		if err != nil {
			panic(err)
		}

		var response strings.Builder

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}

			line = strings.TrimSpace(line)

			if line == "END" {
				fmt.Println("\nSolomon:")
				fmt.Println(response.String())
				break
			}

			response.WriteString(line)
			response.WriteString("\n")
		}
	}

}
