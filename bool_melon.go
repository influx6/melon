package melon

import (
	"github.com/influx6/faux/context"
)

// BoolUniqueHash defines a unique hash for Bool which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const BoolUniqueHash = "189fa523325f8a701fd00ad1b0cd386b4b629299"

// BoolHandler defines a function type receiving both reader and writer types.
type BoolHandler func(BoolReader, BoolWriter) error

// BoolHandlerWithCtx defines a function type receiving both reader and writer types.
type BoolHandlerWithCtx func(context.Context, BoolReader, BoolWriter) error

// BoolStream defines a function type receiving both reader and writer types.
type BoolStream func(BoolStreamReader, BoolStreamWriter) error

// BoolStreamWithCtx defines a function type receiving both reader and writer types.
type BoolStreamWithCtx func(context.Context, BoolStreamReader, BoolStreamWriter) error

// BoolReader defines an interface for reading a single bool type.
type BoolReader interface {
	ReadBool() (bool, error)
}

// BoolReadCloser defines an interface for reading a single bool type.
type BoolReadCloser interface {
	Closer
	BoolReader
}

// BoolStreamReader defines an interface for reading a slice of bool type.
type BoolStreamReader interface {
	Read(int) ([]bool, error)
}

// BoolStreamReadCloser defines an interface for reading a single bool type.
type BoolStreamReadCloser interface {
	Closer
	BoolStreamReader
}

// BoolWriter defines an interface for writing a single bool type.
type BoolWriter interface {
	WriteBool(bool) (int, error)
}

// BoolWriteCloser defines an interface for writing a single bool type.
type BoolWriteCloser interface {
	Closer
	BoolWriter
}

// BoolStreamWriter defines an interface for writing a slice of bool type.
type BoolStreamWriter interface {
	Write([]bool) (int, error)
}

// BoolStreamWriteCloser defines an interface for writing a single bool type.
type BoolStreamWriteCloser interface {
	Closer
	BoolStreamWriter
}
