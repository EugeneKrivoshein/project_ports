package handlers

import (
	"context"
	"testing"
	"time"
)

// Мок-порт для тестирования обработчиков
type MockPort struct {
	id int
}

func (m *MockPort) GetID() int                           { return m.id }
func (m *MockPort) Read() (int, error)                   { return 1, nil }
func (m *MockPort) Write(transactionID, value int) error { return nil }

func TestHandleInPort(t *testing.T) {
	inPort := &MockPort{id: 1}

	// Создаем контекст с отменой, чтобы управлять завершением горутины
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Запускаем обработчик с передачей контекста
	go HandleInPort(ctx, inPort)

	// Даем горутине немного времени на выполнение
	time.Sleep(2 * time.Second)

	// Завершаем работу горутины
	cancel()
}

func TestHandleOutPort(t *testing.T) {
	outPort := &MockPort{id: 2}

	// Создаем контекст с отменой для управления завершением горутины
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Запускаем обработчик OUT порта
	go HandleOutPort(ctx, outPort)

	// Даем горутине немного времени на выполнение
	time.Sleep(2 * time.Second)

	// Завершаем работу горутины
	cancel()
}
