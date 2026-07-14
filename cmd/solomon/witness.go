package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func consultWitness(address, question string) (string, error) {
	conn, err := net.DialTimeout(
		"tcp",
		address,
		2*time.Second,
	)
	if err != nil {
		return "", err
	}
	conn.SetDeadline(
		time.Now().Add(2 * time.Second),
	)
	defer conn.Close()

	message := fmt.Sprintf(
		"TYPE:EVIDENCE\nQUESTION:%s\nEND\n",
		question,
	)

	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", err
	}
	reader := bufio.NewReader(conn)

	var response strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		line = strings.TrimSpace(line)

		if line == "END" {
			testimony := parseTestimony(response.String())
			return testimony.Suspect, nil
		}

		response.WriteString(line)
		response.WriteString("\n")
	}
}
