package melon

import (
	"github.com/influx6/faux/context"
)

// InterfaceUniqueHash defines a unique hash for Interface which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const InterfaceUniqueHash = "4b3ea5d9542287d43b2938e455ad877a8cc1b573"

// InterfaceHandler defines a function type receiving both reader and writer types.
type InterfaceHandler func(InterfaceReader, InterfaceWriter) error

// InterfaceHandlerWithCtx defines a function type receiving both reader and writer types.
type InterfaceHandlerWithCtx func(context.Context, InterfaceReader, InterfaceWriter) error

// InterfaceStream defines a function type receiving both reader and writer types.
type InterfaceStream func(InterfaceStreamReader, InterfaceStreamWriter) error

// InterfaceStreamWithCtx defines a function type receiving both reader and writer types.
type InterfaceStreamWithCtx func(context.Context, InterfaceStreamReader, InterfaceStreamWriter) error

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
