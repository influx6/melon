package melon

import (
	"github.com/influx6/faux/context"
)

// InterfaceUnitReaderFunc defines a function which expects the giving InterfaceReader type has input.
type InterfaceUnitReaderFunc func(InterfaceReader) InterfaceUnitReader

// InterfaceUnitReaderFuncWithContext defines a function which expects the giving InterfaceReader type has input.
// This expects to receive a context.Context type.
type InterfaceUnitReaderFuncWithContext func(context.Context, InterfaceReader) InterfaceUnitReader

// InterfaceReaderFunc defines a function which expects the giving InterfaceReader type has input.
type InterfaceReaderFunc func(InterfaceReader) error

// InterfaceReaderFuncWithContext defines a function which expects the giving InterfaceReader type has input.
// This expects to receive a context.Context type.
type InterfaceReaderFuncWithContext func(context.Context, InterfaceReader) error

// InterfaceReadCloserFunc defines a function which expects the giving InterfaceReadCloser type has input.
type InterfaceReadCloserFunc func(InterfaceReadCloser) error

// InterfaceReadCloserFuncWithContext defines a function which expects the giving InterfaceReadCloser type has input.
// This expects to receive a context.Context type.
type InterfaceReadCloserFuncWithContext func(context.Context, InterfaceReadCloser) error

// InterfaceWriterFunc defines a function which expects the giving InterfaceWriter type has input.
type InterfaceWriterFunc func(InterfaceWriter) error

// InterfaceWriteCloserFunc defines a function which expects the giving InterfaceWriteCloser type has input.
type InterfaceWriteCloserFunc func(InterfaceWriteCloser) error

// InterfaceWriterFuncWithContext defines a function which expects the giving InterfaceWriter type has input.
// This expects to receive a context.Context type.
type InterfaceWriterFuncWithContext func(context.Context, InterfaceWriter) error

// InterfaceWriteCloserFuncWithContext defines a function which expects the giving InterfaceWriteCloser type has input.
// This expects to receive a context.Context type.
type InterfaceWriteCloserFuncWithContext func(context.Context, InterfaceWriteCloser) error

// InterfaceReader defines an interface for reading a slice of interface{} types.
type InterfaceReader interface {
	Read([]interface{}) (int, error)
}

// InterfaceReadCloser defines an interface for reading a slice of interface{} types.
type InterfaceReadCloser interface {
	Closer
	InterfaceReader
}

// InterfaceUnitReader defines an interface for reading a single item of interface{} type.
type InterfaceUnitReader interface {
	ReadUnit() (interface{}, error)
}

// InterfaceWriter defines an interface for writing a slice of interface{} types.
type InterfaceWriter interface {
	Write([]interface{}) (int, error)
}

// InterfaceUnitWriter defines an interface for writing a single interface{} type.
type InterfaceUnitWriter interface {
	WriteUnit(interface{}) (int, error)
}

// InterfaceWriteCloser defines an interface for writing a slice of interface{} types.
type InterfaceWriteCloser interface {
	Closer
	InterfaceWriter
}
