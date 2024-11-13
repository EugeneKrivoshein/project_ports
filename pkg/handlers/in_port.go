package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/EugeneKrivoshein/project_ports/pkg/ports"
)

// HandleInPort запускает горутину для IN порта, выполняя периодическое чтение
func HandleInPort(ctx context.Context, port ports.ReadablePort) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("IN Port %d stopping...\n", port.GetID())
			return
		default:
			value, err := port.Read()
			if err != nil {
				fmt.Println("Error reading from IN port:", err)
				return
			}
			log.Printf("Read value %d from IN port %d\n", value, port.GetID())
			time.Sleep(1 * time.Second)
		}
	}
}
