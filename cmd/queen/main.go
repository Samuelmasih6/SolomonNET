package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
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
		var message string
		fmt.Print(
			"Challenge Type (RIDDLE/CASE/LOOKUP/HISTORY): ",
		)
		challengetype, err := consolereader.ReadString('\n')
		if err != nil {
			return
		}
		challengetype = strings.TrimSpace(challengetype)
		challengetype =
			strings.ToUpper(
				strings.TrimSpace(challengetype),
			)
		if challengetype == "EXIT" {
			return
		}
		if challengetype == "HISTORY" {

			message = "TYPE:HISTORY\nEND\n"

		} else {
			if challengetype == "LOOKUP" {
				fmt.Print("CASE ID:")
			} else {
				fmt.Print("Question:")
			}
			question, err := consolereader.ReadString('\n')
			if err != nil {
				return
			}

			question = strings.TrimSpace(question)

			if question == "exit" {
				return
			}

			if challengetype == "LOOKUP" {

				id, err := strconv.Atoi(question)
				if err != nil {
					fmt.Println("Invalid case ID")
					continue
				}

				message = fmt.Sprintf(
					"TYPE:LOOKUP\nCASE_ID:%d\nEND\n",
					id,
				)

			} else {

				message = fmt.Sprintf(
					"TYPE:%s\nQUESTION:%s\nEND\n",
					challengetype,
					question,
				)
			}
		}
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
