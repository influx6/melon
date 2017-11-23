package melon

import (
	"github.com/influx6/faux/context"
)

// Float64UniqueHash defines a unique hash for Float64 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Float64UniqueHash = "c11c63ba6f4639c386f2b6ed2ec7dec76ff35a5d"

// Float64Handler defines a function type receiving both reader and writer types.
type Float64Handler func(Float64Reader, Float64Writer) error

// Float64HandlerWithCtx defines a function type receiving both reader and writer types.
type Float64HandlerWithCtx func(context.Context, Float64Reader, Float64Writer) error

// Float64Stream defines a function type receiving both reader and writer types.
type Float64Stream func(Float64StreamReader, Float64StreamWriter) error

// Float64StreamWithCtx defines a function type receiving both reader and writer types.
type Float64StreamWithCtx func(context.Context, Float64StreamReader, Float64StreamWriter) error

// Float64Reader defines an interface for reading a single float64 type.
type Float64Reader interface {
	ReadFloat64() (float64, error)
}

// Float64ReadCloser defines an interface for reading a single float64 type.
type Float64ReadCloser interface {
	Closer
	Float64Reader
}

// Float64StreamReader defines an interface for reading a slice of float64 type.
type Float64StreamReader interface {
	Read(int) ([]float64, error)
}

// Float64StreamReadCloser defines an interface for reading a single float64 type.
type Float64StreamReadCloser interface {
	Closer
	Float64StreamReader
}

// Float64Writer defines an interface for writing a single float64 type.
type Float64Writer interface {
	WriteFloat64(float64) (int, error)
}

// Float64WriteCloser defines an interface for writing a single float64 type.
type Float64WriteCloser interface {
	Closer
	Float64Writer
}

// Float64StreamWriter defines an interface for writing a slice of float64 type.
type Float64StreamWriter interface {
	Write([]float64) (int, error)
}

// Float64StreamWriteCloser defines an interface for writing a single float64 type.
type Float64StreamWriteCloser interface {
	Closer
	Float64StreamWriter
}
