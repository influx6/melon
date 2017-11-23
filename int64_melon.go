package melon

import (
	"github.com/influx6/faux/context"
)

// Int64UniqueHash defines a unique hash for Int64 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Int64UniqueHash = "d9735f419182b6c737eba2e4b2d3f05a5b349bc9"

// Int64Handler defines a function type receiving both reader and writer types.
type Int64Handler func(Int64Reader, Int64Writer) error

// Int64HandlerWithCtx defines a function type receiving both reader and writer types.
type Int64HandlerWithCtx func(context.Context, Int64Reader, Int64Writer) error

// Int64Stream defines a function type receiving both reader and writer types.
type Int64Stream func(Int64StreamReader, Int64StreamWriter) error

// Int64StreamWithCtx defines a function type receiving both reader and writer types.
type Int64StreamWithCtx func(context.Context, Int64StreamReader, Int64StreamWriter) error

// Int64Reader defines an interface for reading a single int64 type.
type Int64Reader interface {
	ReadInt64() (int64, error)
}

// Int64ReadCloser defines an interface for reading a single int64 type.
type Int64ReadCloser interface {
	Closer
	Int64Reader
}

// Int64StreamReader defines an interface for reading a slice of int64 type.
type Int64StreamReader interface {
	Read(int) ([]int64, error)
}

// Int64StreamReadCloser defines an interface for reading a single int64 type.
type Int64StreamReadCloser interface {
	Closer
	Int64StreamReader
}

// Int64Writer defines an interface for writing a single int64 type.
type Int64Writer interface {
	WriteInt64(int64) (int, error)
}

// Int64WriteCloser defines an interface for writing a single int64 type.
type Int64WriteCloser interface {
	Closer
	Int64Writer
}

// Int64StreamWriter defines an interface for writing a slice of int64 type.
type Int64StreamWriter interface {
	Write([]int64) (int, error)
}

// Int64StreamWriteCloser defines an interface for writing a single int64 type.
type Int64StreamWriteCloser interface {
	Closer
	Int64StreamWriter
}
