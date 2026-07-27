package main

import "sync"

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
