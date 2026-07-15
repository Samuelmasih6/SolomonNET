package main

import (
	"sync"
	"time"
)

type Court struct {
	mu     sync.RWMutex
	nextID int
	cases  map[int]Case
}

func NewCourt() *Court {
	return &Court{
		nextID: 1,
		cases:  make(map[int]Case),
	}
}

func (c *Court) CreateCase(
	question string,
	testimonies []string,
	verdict string,
	confidence string,
) Case {

	c.mu.Lock()
	defer c.mu.Unlock()

	id := c.nextID
	c.nextID++

	newCase := Case{
		ID:          id,
		Question:    question,
		Testimonies: testimonies,
		Verdict:     verdict,
		Confidence:  confidence,
		CreatedAt:   time.Now(),
	}

	c.cases[id] = newCase

	return newCase
}

func (c *Court) GetCase(id int) (Case, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cse, ok := c.cases[id]
	return cse, ok
}

func (c *Court) ListCases() []Case {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cases := make([]Case, 0, len(c.cases))

	for _, cse := range c.cases {
		cases = append(cases, cse)
	}

	return cases
}
