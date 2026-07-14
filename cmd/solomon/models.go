package main

type Challenge struct {
	Type     string
	Question string
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
	ID         int
	Question   string
	Verdict    string
	Confidence string
}
