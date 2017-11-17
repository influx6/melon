package melon

import (
	"github.com/influx6/faux/context"
)

// Int16UnitReaderFunc defines a function which expects the giving Int16Reader type has input.
type Int16UnitReaderFunc func(Int16Reader) Int16UnitReader

// Int16UnitReaderFuncWithContext defines a function which expects the giving Int16Reader type has input.
// This expects to receive a context.Context type.
type Int16UnitReaderFuncWithContext func(context.Context, Int16Reader) Int16UnitReader

// Int16ReaderFunc defines a function which expects the giving Int16Reader type has input.
type Int16ReaderFunc func(Int16Reader) error

// Int16ReaderFuncWithContext defines a function which expects the giving Int16Reader type has input.
// This expects to receive a context.Context type.
type Int16ReaderFuncWithContext func(context.Context, Int16Reader) error

// Int16ReadCloserFunc defines a function which expects the giving Int16ReadCloser type has input.
type Int16ReadCloserFunc func(Int16ReadCloser) error

// Int16ReadCloserFuncWithContext defines a function which expects the giving Int16ReadCloser type has input.
// This expects to receive a context.Context type.
type Int16ReadCloserFuncWithContext func(context.Context, Int16ReadCloser) error

// Int16WriterFunc defines a function which expects the giving Int16Writer type has input.
type Int16WriterFunc func(Int16Writer) error

// Int16WriteCloserFunc defines a function which expects the giving Int16WriteCloser type has input.
type Int16WriteCloserFunc func(Int16WriteCloser) error

// Int16WriterFuncWithContext defines a function which expects the giving Int16Writer type has input.
// This expects to receive a context.Context type.
type Int16WriterFuncWithContext func(context.Context, Int16Writer) error

// Int16WriteCloserFuncWithContext defines a function which expects the giving Int16WriteCloser type has input.
// This expects to receive a context.Context type.
type Int16WriteCloserFuncWithContext func(context.Context, Int16WriteCloser) error

// Int16Reader defines an interface for reading a slice of int16 types.
type Int16Reader interface {
	Read([]int16) (int, error)
}

// Int16ReadCloser defines an interface for reading a slice of int16 types.
type Int16ReadCloser interface {
	Closer
	Int16Reader
}

// Int16UnitReader defines an interface for reading a single item of int16 type.
type Int16UnitReader interface {
	Read() (int16, error)
}

// Int16Writer defines an interface for writing a slice of int16 types.
type Int16Writer interface {
	Write([]int16) (int, error)
}

// Int16WriteCloser defines an interface for writing a slice of int16 types.
type Int16WriteCloser interface {
	Closer
	Int16Writer
}
