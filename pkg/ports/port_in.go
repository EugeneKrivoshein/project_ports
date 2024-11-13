package ports

import "math/rand"

type InPort struct {
	ID int
}

func (p *InPort) GetID() int {
	return p.ID
}

func (p *InPort) Read() (int, error) {
	// Логика чтения значения (например, случайное число 0 или 1)
	return rand.Intn(2), nil
}
