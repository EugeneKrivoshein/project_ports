package config

import (
	"context"
	"log"
	"os"
	"strconv"
)

// Config содержит количество IN и OUT портов и контекст завершения
type Config struct {
	NumInPorts  int
	NumOutPorts int
	Ctx         context.Context
	Cancel      context.CancelFunc
}

// LoadConfig загружает конфигурацию и создает контекст завершения
func LoadConfig() *Config {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <num_in_ports> <num_out_ports>\n", os.Args[0])
	}

	numInPorts, err := strconv.Atoi(os.Args[1])
	if err != nil || numInPorts < 0 {
		log.Fatalf("Invalid number of IN ports: %v\n", os.Args[1])
	}

	numOutPorts, err := strconv.Atoi(os.Args[2])
	if err != nil || numOutPorts < 0 {
		log.Fatalf("Invalid number of OUT ports: %v\n", os.Args[2])
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Config{
		NumInPorts:  numInPorts,
		NumOutPorts: numOutPorts,
		Ctx:         ctx,
		Cancel:      cancel,
	}
}
