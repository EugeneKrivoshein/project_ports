package api

import (
	"fmt"

	"github.com/EugeneKrivoshein/project_ports/pkg/ports"
)

// ReadPort читает значение из IN порта
func ReadPort(port ports.ReadablePort) {
	value, err := port.Read()
	if err != nil {
		fmt.Printf("Error reading from port %d: %v\n", port.GetID(), err)
		return
	}
	fmt.Printf("IN Port %d: Read Value %d\n", port.GetID(), value)
}

// WritePort записывает значение в OUT порт
func WritePort(port ports.WritablePort, transactionID, value int) {
	err := port.Write(transactionID, value)
	if err != nil {
		fmt.Printf("Error writing to port %d: %v\n", port.GetID(), err)
		return
	}
}
