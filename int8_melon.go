package melon

import (
	"github.com/influx6/faux/context"
)

// Int8UnitReaderFunc defines a function which expects the giving Int8Reader type has input.
type Int8UnitReaderFunc func(Int8Reader) Int8UnitReader

// Int8UnitReaderFuncWithContext defines a function which expects the giving Int8Reader type has input.
// This expects to receive a context.Context type.
type Int8UnitReaderFuncWithContext func(context.Context, Int8Reader) Int8UnitReader

// Int8ReaderFunc defines a function which expects the giving Int8Reader type has input.
type Int8ReaderFunc func(Int8Reader) error

// Int8ReaderFuncWithContext defines a function which expects the giving Int8Reader type has input.
// This expects to receive a context.Context type.
type Int8ReaderFuncWithContext func(context.Context, Int8Reader) error

// Int8ReadCloserFunc defines a function which expects the giving Int8ReadCloser type has input.
type Int8ReadCloserFunc func(Int8ReadCloser) error

// Int8ReadCloserFuncWithContext defines a function which expects the giving Int8ReadCloser type has input.
// This expects to receive a context.Context type.
type Int8ReadCloserFuncWithContext func(context.Context, Int8ReadCloser) error

// Int8WriterFunc defines a function which expects the giving Int8Writer type has input.
type Int8WriterFunc func(Int8Writer) error

// Int8WriteCloserFunc defines a function which expects the giving Int8WriteCloser type has input.
type Int8WriteCloserFunc func(Int8WriteCloser) error

// Int8WriterFuncWithContext defines a function which expects the giving Int8Writer type has input.
// This expects to receive a context.Context type.
type Int8WriterFuncWithContext func(context.Context, Int8Writer) error

// Int8WriteCloserFuncWithContext defines a function which expects the giving Int8WriteCloser type has input.
// This expects to receive a context.Context type.
type Int8WriteCloserFuncWithContext func(context.Context, Int8WriteCloser) error

// Int8Reader defines an interface for reading a slice of int8 types.
type Int8Reader interface {
	Read([]int8) (int, error)
}

// Int8ReadCloser defines an interface for reading a slice of int8 types.
type Int8ReadCloser interface {
	Closer
	Int8Reader
}

// Int8UnitReader defines an interface for reading a single item of int8 type.
type Int8UnitReader interface {
	Read() (int8, error)
}

// Int8Writer defines an interface for writing a slice of int8 types.
type Int8Writer interface {
	Write([]int8) (int, error)
}

// Int8WriteCloser defines an interface for writing a slice of int8 types.
type Int8WriteCloser interface {
	Closer
	Int8Writer
}
