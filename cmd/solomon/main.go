package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type Challenge struct {
	Type     string
	Question string
}

type Testimony struct {
	Type   string
	Answer string
}

type WitnessResult struct {
	Address   string
	Testimony string
	Err       error
}

func handleConnection(conn net.Conn, id int) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	var message strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("[Queen %d] disconnected\n", id)
			return
		}

		line = strings.TrimSpace(line)

		if line == "END" {
			challenge := parseChallenge(message.String())

			fmt.Printf(
				"[Queen %d] Type=%s Question=%s\n",
				id,
				challenge.Type,
				challenge.Question,
			)

			var answer string

			switch challenge.Type {

			case "RIDDLE":
				answer = solveRiddle(challenge.Question)

			case "CASE":
				witnesses := []string{
					"localhost:9091",
					"localhost:9092",
					"localhost:9093",
				}
				results := make(chan WitnessResult)

				for _, address := range witnesses {

					go func(addr string) {

						testimony, err := consultWitness(
							addr,
							challenge.Question,
						)

						results <- WitnessResult{
							Address:   addr,
							Testimony: testimony,
							Err:       err,
						}

					}(address)
				}

				var testimonies []string

				for i := 0; i < len(witnesses); i++ {

					result := <-results

					if result.Err != nil {
						testimonies = append(
							testimonies,
							fmt.Sprintf(
								"%s -> unavailable",
								result.Address,
							),
						)
					} else {
						testimonies = append(
							testimonies,
							fmt.Sprintf(
								"%s -> %s",
								result.Address,
								result.Testimony,
							),
						)
					}
				}

				answer = strings.Join(
					testimonies,
					"\n",
				)

			default:
				answer = "Unknown challenge type"
			}

			response := fmt.Sprintf(
				"TYPE:ANSWER\nANSWER:%s\nEND\n",
				answer,
			)

			_, err = conn.Write([]byte(response))
			if err != nil {
				fmt.Printf(
					"[Queen %d] failed to send response: %v\n",
					id,
					err,
				)
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

func parseTestimony(message string) Testimony {
	var msgType string
	var answer string

	lines := strings.Split(message, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "TYPE:") {
			msgType = strings.TrimPrefix(line, "TYPE:")
		}

		if strings.HasPrefix(line, "ANSWER:") {
			answer = strings.TrimPrefix(line, "ANSWER:")
		}
	}

	return Testimony{
		Type:   msgType,
		Answer: answer,
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
			return testimony.Answer, nil
		}

		response.WriteString(line)
		response.WriteString("\n")
	}
}

func main() {
	var nextID int

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

		nextID++
		id := nextID

		fmt.Printf("Queen %d connected\n", id)

		go handleConnection(conn, id)
	}
}
