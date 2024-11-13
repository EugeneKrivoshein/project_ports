package ports

type PortType int

const (
	IN PortType = iota
	OUT
)

type Port interface {
	GetID() int
}

type ReadablePort interface {
	Port
	Read() (int, error)
}

type WritablePort interface {
	Port
	Write(transactionID, value int) error
}

func NewPort(id int, portType PortType) interface{} {
	switch portType {
	case IN:
		return &InPort{ID: id}
	case OUT:
		return &OutPort{ID: id}
	default:
		return nil
	}
}
