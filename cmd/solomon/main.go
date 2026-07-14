package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

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
				votes := make(map[string]int)
				for _, address := range witnesses {

					go func(addr string) {

						testimony, err := consultWitness(
							addr,
							challenge.Question,
						)

						results <- WitnessResult{
							Address: addr,
							Suspect: testimony,
							Err:     err,
						}

					}(address)
				}

				var testimonies []string
				availableWitnesses := 0

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
						availableWitnesses++
						testimonies = append(
							testimonies,
							fmt.Sprintf(
								"%s -> %s",
								result.Address,
								result.Suspect,
							),
						)
						votes[result.Suspect]++
					}
				}

				var winner string
				var maxVotes int

				for suspect, count := range votes {

					if count > maxVotes {
						maxVotes = count
						winner = suspect
					}
				}
				confidence := fmt.Sprintf(
					"%d/%d",
					maxVotes,
					availableWitnesses,
				)
				if maxVotes <= availableWitnesses/2 {
					winner = "INCONCLUSIVE"
				}
				answer = fmt.Sprintf(
					"VERDICT:%s\nCONFIDENCE:%s",
					winner,
					confidence,
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
