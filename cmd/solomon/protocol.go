package main

import (
	"strconv"
	"strings"
)

func parseChallenge(message string) Challenge {
	var msgType string
	var question string
	var caseID int

	lines := strings.Split(message, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "TYPE:") {
			msgType = strings.TrimPrefix(line, "TYPE:")
		}

		if strings.HasPrefix(line, "QUESTION:") {
			question = strings.TrimPrefix(line, "QUESTION:")
		}

		if strings.HasPrefix(line, "CASE_ID:") {
			idStr := strings.TrimPrefix(
				line,
				"CASE_ID:",
			)
			id, err := strconv.Atoi(idStr)
			if err == nil {
				caseID = id
			}
		}
	}

	return Challenge{
		Type:     msgType,
		Question: question,
		CaseID:   caseID,
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
