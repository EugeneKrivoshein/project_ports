package ports

import (
	"testing"
)

func TestInPortRead(t *testing.T) {
	inPort := NewPort(1, IN).(*InPort)

	value, err := inPort.Read()
	if err != nil || (value != 0 && value != 1) {
		t.Errorf("Expected value 0 or 1, got %d with error %v", value, err)
	}
}

func TestOutPortWrite(t *testing.T) {
	outPort := NewPort(2, OUT).(*OutPort)

	err := outPort.Write(100, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
