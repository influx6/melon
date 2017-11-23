package melon

import (
	"github.com/influx6/faux/context"
)

// Int32UniqueHash defines a unique hash for Int32 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Int32UniqueHash = "f90a5efe8610965bee01abba15d0e1f54046e2e8"

// Int32Handler defines a function type receiving both reader and writer types.
type Int32Handler func(Int32Reader, Int32Writer) error

// Int32HandlerWithCtx defines a function type receiving both reader and writer types.
type Int32HandlerWithCtx func(context.Context, Int32Reader, Int32Writer) error

// Int32Stream defines a function type receiving both reader and writer types.
type Int32Stream func(Int32StreamReader, Int32StreamWriter) error

// Int32StreamWithCtx defines a function type receiving both reader and writer types.
type Int32StreamWithCtx func(context.Context, Int32StreamReader, Int32StreamWriter) error

// Int32Reader defines an interface for reading a single int32 type.
type Int32Reader interface {
	ReadInt32() (int32, error)
}

// Int32ReadCloser defines an interface for reading a single int32 type.
type Int32ReadCloser interface {
	Closer
	Int32Reader
}

// Int32StreamReader defines an interface for reading a slice of int32 type.
type Int32StreamReader interface {
	Read(int) ([]int32, error)
}

// Int32StreamReadCloser defines an interface for reading a single int32 type.
type Int32StreamReadCloser interface {
	Closer
	Int32StreamReader
}

// Int32Writer defines an interface for writing a single int32 type.
type Int32Writer interface {
	WriteInt32(int32) (int, error)
}

// Int32WriteCloser defines an interface for writing a single int32 type.
type Int32WriteCloser interface {
	Closer
	Int32Writer
}

// Int32StreamWriter defines an interface for writing a slice of int32 type.
type Int32StreamWriter interface {
	Write([]int32) (int, error)
}

// Int32StreamWriteCloser defines an interface for writing a single int32 type.
type Int32StreamWriteCloser interface {
	Closer
	Int32StreamWriter
}
