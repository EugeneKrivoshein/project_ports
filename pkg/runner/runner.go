package runner

import (
	"context"
	"sync"

	"github.com/EugeneKrivoshein/project_ports/pkg/api"
	"github.com/EugeneKrivoshein/project_ports/pkg/handlers"
	"github.com/EugeneKrivoshein/project_ports/pkg/ports"
)

func RunPorts(numInPorts, numOutPorts int, ctx context.Context, wg *sync.WaitGroup) ([]ports.ReadablePort, []ports.WritablePort) {
	inPorts := make([]ports.ReadablePort, numInPorts)
	outPorts := make([]ports.WritablePort, numOutPorts)

	for i := 0; i < numInPorts; i++ {
		inPort := &ports.InPort{ID: i}
		inPorts[i] = inPort
		wg.Add(1)
		go func(port ports.ReadablePort) {
			defer wg.Done()
			handlers.HandleInPort(ctx, port)
		}(inPort)
	}

	for i := 0; i < numOutPorts; i++ {
		outPort := &ports.OutPort{ID: i}
		outPorts[i] = outPort
		wg.Add(1)
		go func(port ports.WritablePort) {
			defer wg.Done()
			handlers.HandleOutPort(ctx, port)
		}(outPort)
	}

	return inPorts, outPorts
}

func TestAPI(inPorts []ports.ReadablePort, outPorts []ports.WritablePort) {
	for _, port := range inPorts {
		api.ReadPort(port)
	}
	for i, port := range outPorts {
		api.WritePort(port, i, i%2)
	}
}
