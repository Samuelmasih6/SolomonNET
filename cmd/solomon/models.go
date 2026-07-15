package main

import "time"

type Challenge struct {
	Type     string
	Question string
	CaseID   int
}

type Testimony struct {
	Type    string
	Suspect string
}

type WitnessResult struct {
	Address string
	Suspect string
	Err     error
}

type Case struct {
	ID          int
	Question    string
	Testimonies []string
	Verdict     string
	Confidence  string
	CreatedAt   time.Time
}
