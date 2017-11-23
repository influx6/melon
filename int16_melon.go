package melon

import (
	"github.com/influx6/faux/context"
)

// Int16UniqueHash defines a unique hash for Int16 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Int16UniqueHash = "3b1227b09fd0719b0d9a553739bfa14479fc0a09"

// Int16Handler defines a function type receiving both reader and writer types.
type Int16Handler func(Int16Reader, Int16Writer) error

// Int16HandlerWithCtx defines a function type receiving both reader and writer types.
type Int16HandlerWithCtx func(context.Context, Int16Reader, Int16Writer) error

// Int16Stream defines a function type receiving both reader and writer types.
type Int16Stream func(Int16StreamReader, Int16StreamWriter) error

// Int16StreamWithCtx defines a function type receiving both reader and writer types.
type Int16StreamWithCtx func(context.Context, Int16StreamReader, Int16StreamWriter) error

// Int16Reader defines an interface for reading a single int16 type.
type Int16Reader interface {
	ReadInt16() (int16, error)
}

// Int16ReadCloser defines an interface for reading a single int16 type.
type Int16ReadCloser interface {
	Closer
	Int16Reader
}

// Int16StreamReader defines an interface for reading a slice of int16 type.
type Int16StreamReader interface {
	Read(int) ([]int16, error)
}

// Int16StreamReadCloser defines an interface for reading a single int16 type.
type Int16StreamReadCloser interface {
	Closer
	Int16StreamReader
}

// Int16Writer defines an interface for writing a single int16 type.
type Int16Writer interface {
	WriteInt16(int16) (int, error)
}

// Int16WriteCloser defines an interface for writing a single int16 type.
type Int16WriteCloser interface {
	Closer
	Int16Writer
}

// Int16StreamWriter defines an interface for writing a slice of int16 type.
type Int16StreamWriter interface {
	Write([]int16) (int, error)
}

// Int16StreamWriteCloser defines an interface for writing a single int16 type.
type Int16StreamWriteCloser interface {
	Closer
	Int16StreamWriter
}
