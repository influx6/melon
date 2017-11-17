package melon

import (
	"github.com/influx6/faux/context"
)

// Int32ReaderFunc defines a function which expects the giving Int32Reader type has input.
type Int32ReaderFunc func(Int32Reader) error

// Int32ReaderFuncWithContext defines a function which expects the giving Int32Reader type has input.
// This expects to receive a context.Context type.
type Int32ReaderFuncWithContext func(context.Context, Int32Reader) error

// Int32ReadCloserFunc defines a function which expects the giving Int32ReadCloser type has input.
type Int32ReadCloserFunc func(Int32ReadCloser) error

// Int32ReadCloserFuncWithContext defines a function which expects the giving Int32ReadCloser type has input.
// This expects to receive a context.Context type.
type Int32ReadCloserFuncWithContext func(context.Context, Int32ReadCloser) error

// Int32WriterFunc defines a function which expects the giving Int32Writer type has input.
type Int32WriterFunc func(Int32Writer) error

// Int32WriteCloserFunc defines a function which expects the giving Int32WriteCloser type has input.
type Int32WriteCloserFunc func(Int32WriteCloser) error

// Int32WriterFuncWithContext defines a function which expects the giving Int32Writer type has input.
// This expects to receive a context.Context type.
type Int32WriterFuncWithContext func(context.Context, Int32Writer) error

// Int32WriteCloserFuncWithContext defines a function which expects the giving Int32WriteCloser type has input.
// This expects to receive a context.Context type.
type Int32WriteCloserFuncWithContext func(context.Context, Int32WriteCloser) error

// Int32Reader defines an interface for reading a slice of int32 types.
type Int32Reader interface {
	Read([]int32) (int, error)
}

// Int32ReadCloser defines an interface for reading a slice of int32 types.
type Int32ReadCloser interface {
	Closer
	Int32Reader
}

// Int32Writer defines an interface for writing a slice of int32 types.
type Int32Writer interface {
	Write([]int32) (int, error)
}

// Int32WriteCloser defines an interface for writing a slice of int32 types.
type Int32WriteCloser interface {
	Closer
	Int32Writer
}
