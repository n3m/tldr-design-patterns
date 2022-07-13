package logistics

/*
	HOW-TO: Factory Method Pattern
*/

/*
	CONTEXT:
	In this specific case, we're using Go's interfaces in order to generate an Object-Oriented implementation hierarchy
*/

// This interface will represent the father of the "category" of objects that we want to generate
type ILogisticTransport interface {
	Deliver()
	Load()
}

// This struct will generate all the baseline implementations of the interface, and we will derive all
// of our other objects from this
type transport struct {
}

func (t *transport) Deliver() {

}

func (t *transport) Load() {

}

// In the following lines we're deriving a Bus and a Ship structure that inherits the transport structure properties and that also implements the ILogisticTransport interface
type Bus struct {
	transport
}

type Ship struct {
	transport
}

// This method will be in charge of being our "Creator"
// Basically we just need to pass some kind of type to this method,
// and it will return a new instance of an object that we want to generate
// that meets the interface requirements
func CreateNewTransport(t string) ILogisticTransport {
	switch t {
	case "bus":
		return &Bus{}
	case "ship":
		return &Ship{}
	default:
		return nil
	}

}
