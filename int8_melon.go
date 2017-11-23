package melon

import (
	"github.com/influx6/faux/context"
)

// Int8UniqueHash defines a unique hash for Int8 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Int8UniqueHash = "da945acc637aec1ea9d814cc20b2ed6e794ef8b1"

// Int8Handler defines a function type receiving both reader and writer types.
type Int8Handler func(Int8Reader, Int8Writer) error

// Int8HandlerWithCtx defines a function type receiving both reader and writer types.
type Int8HandlerWithCtx func(context.Context, Int8Reader, Int8Writer) error

// Int8Stream defines a function type receiving both reader and writer types.
type Int8Stream func(Int8StreamReader, Int8StreamWriter) error

// Int8StreamWithCtx defines a function type receiving both reader and writer types.
type Int8StreamWithCtx func(context.Context, Int8StreamReader, Int8StreamWriter) error

// Int8Reader defines an interface for reading a single int8 type.
type Int8Reader interface {
	ReadInt8() (int8, error)
}

// Int8ReadCloser defines an interface for reading a single int8 type.
type Int8ReadCloser interface {
	Closer
	Int8Reader
}

// Int8StreamReader defines an interface for reading a slice of int8 type.
type Int8StreamReader interface {
	Read(int) ([]int8, error)
}

// Int8StreamReadCloser defines an interface for reading a single int8 type.
type Int8StreamReadCloser interface {
	Closer
	Int8StreamReader
}

// Int8Writer defines an interface for writing a single int8 type.
type Int8Writer interface {
	WriteInt8(int8) (int, error)
}

// Int8WriteCloser defines an interface for writing a single int8 type.
type Int8WriteCloser interface {
	Closer
	Int8Writer
}

// Int8StreamWriter defines an interface for writing a slice of int8 type.
type Int8StreamWriter interface {
	Write([]int8) (int, error)
}

// Int8StreamWriteCloser defines an interface for writing a single int8 type.
type Int8StreamWriteCloser interface {
	Closer
	Int8StreamWriter
}
