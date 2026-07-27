package main

import (
	"sync"
	"time"
)

type MemoryRepository struct {
	mu     sync.RWMutex
	nextID int
	cases  map[int]Case
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		nextID: 1,
		cases:  make(map[int]Case),
	}
}

func (r *MemoryRepository) Create(
	question string,
	testimonies []string,
	verdict string,
	confidence string,
) Case {

	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.nextID
	r.nextID++

	newCase := Case{
		ID:          id,
		Question:    question,
		Testimonies: testimonies,
		Verdict:     verdict,
		Confidence:  confidence,
		CreatedAt:   time.Now(),
	}

	r.cases[id] = newCase

	return newCase
}
