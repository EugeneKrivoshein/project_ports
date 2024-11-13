package ports

import "fmt"

type OutPort struct {
	ID int
}

func (p *OutPort) GetID() int {
	return p.ID
}

func (p *OutPort) Write(transactionID, value int) error {
	// запись значения в OUT порт
	fmt.Printf("Transaction %d: Writing value %d to OUT port %d\n", transactionID, value, p.ID)
	return nil
}
