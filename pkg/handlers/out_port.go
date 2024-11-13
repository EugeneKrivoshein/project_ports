package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/EugeneKrivoshein/project_ports/pkg/ports"
)

// HandleOutPort запускает горутину для OUT порта, выполняя запись с задержкой
func HandleOutPort(ctx context.Context, port ports.WritablePort) {
	transactionID := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("OUT Port %d stopping...\n", port.GetID())
			return
		default:
			value := transactionID % 2
			port.Write(transactionID, value)
			transactionID++
			time.Sleep(1 * time.Second)
		}
	}
}
