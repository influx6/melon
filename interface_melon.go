package melon

// InterfaceUniqueHash defines a unique hash for Interface which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const InterfaceUniqueHash = "4b3ea5d9542287d43b2938e455ad877a8cc1b573"

// InterfaceReader defines an interface for reading a single interface{} type.
type InterfaceReader interface {
	ReadInterface() (interface{}, error)
}

// InterfaceReadCloser defines an interface for reading a single interface{} type.
type InterfaceReadCloser interface {
	Closer
	InterfaceReader
}

// InterfaceStreamReader defines an interface for reading a slice of interface{} type.
type InterfaceStreamReader interface {
	Read(int) ([]interface{}, error)
}

// InterfaceStreamReadCloser defines an interface for reading a single interface{} type.
type InterfaceStreamReadCloser interface {
	Closer
	InterfaceStreamReader
}

// InterfaceWriter defines an interface for writing a single interface{} type.
type InterfaceWriter interface {
	WriteInterface(interface{}) (int, error)
}

// InterfaceWriteCloser defines an interface for writing a single interface{} type.
type InterfaceWriteCloser interface {
	Closer
	InterfaceWriter
}

// InterfaceStreamWriter defines an interface for writing a slice of interface{} type.
type InterfaceStreamWriter interface {
	Write([]interface{}) (int, error)
}

// InterfaceStreamWriteCloser defines an interface for writing a single interface{} type.
type InterfaceStreamWriteCloser interface {
	Closer
	InterfaceStreamWriter
}
