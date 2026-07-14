package main

import "strings"

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
		Type:    msgType,
		Suspect: answer,
	}
}
