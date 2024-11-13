package main

import (
	"fmt"
	"sync"

	"github.com/EugeneKrivoshein/project_ports/config"
	"github.com/EugeneKrivoshein/project_ports/pkg/runner"
)

func main() {
	cfg := config.LoadConfig()
	defer cfg.Cancel()

	fmt.Printf("Starting application with %d IN ports and %d OUT ports\n", cfg.NumInPorts, cfg.NumOutPorts)

	var wg sync.WaitGroup

	inPorts, outPorts := runner.RunPorts(cfg.NumInPorts, cfg.NumOutPorts, cfg.Ctx, &wg)

	go runner.TestAPI(inPorts, outPorts)

	wg.Wait()
	fmt.Println("Application stopped.")
}
