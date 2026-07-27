package main

import (
	"sort"
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

func (r *MemoryRepository) Get(id int) (Case, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cse, ok := r.cases[id]
	return cse, ok
}

func (r *MemoryRepository) List() []Case {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cases := make([]Case, 0, len(r.cases))

	for _, cse := range r.cases {
		cases = append(cases, cse)
	}

	sort.Slice(
		cases,
		func(i, j int) bool {
			return cases[i].ID < cases[j].ID
		},
	)

	return cases
}
