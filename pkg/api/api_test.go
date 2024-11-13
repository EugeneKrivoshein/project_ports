package api

import (
	"fmt"
	"testing"
)

// Мок-порты для тестов
type MockInPort struct {
	id int
}

func (m *MockInPort) GetID() int                           { return m.id }
func (m *MockInPort) Read() (int, error)                   { return 1, nil }
func (m *MockInPort) Write(transactionID, value int) error { return nil }

type MockOutPort struct {
	id int
}

func (m *MockOutPort) GetID() int                           { return m.id }
func (m *MockOutPort) Read() (int, error)                   { return 0, fmt.Errorf("unsupported") }
func (m *MockOutPort) Write(transactionID, value int) error { return nil }

func TestReadPort(t *testing.T) {
	inPort := &MockInPort{id: 1}
	ReadPort(inPort) // Проверка на чтение
}

func TestWritePort(t *testing.T) {
	outPort := &MockOutPort{id: 2}
	WritePort(outPort, 100, 1) // Проверка на запись
}
