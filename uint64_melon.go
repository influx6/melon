package melon

import (
	"github.com/influx6/faux/context"
)

// UInt64UniqueHash defines a unique hash for UInt64 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const UInt64UniqueHash = "9dfe93d26e38e94f8348337651317c3130bb021d"

// UInt64Handler defines a function type receiving both reader and writer types.
type UInt64Handler func(UInt64Reader, UInt64Writer) error

// UInt64HandlerWithCtx defines a function type receiving both reader and writer types.
type UInt64HandlerWithCtx func(context.Context, UInt64Reader, UInt64Writer) error

// UInt64Stream defines a function type receiving both reader and writer types.
type UInt64Stream func(UInt64StreamReader, UInt64StreamWriter) error

// UInt64StreamWithCtx defines a function type receiving both reader and writer types.
type UInt64StreamWithCtx func(context.Context, UInt64StreamReader, UInt64StreamWriter) error

// UInt64Reader defines an interface for reading a single uint64 type.
type UInt64Reader interface {
	ReadUInt64() (uint64, error)
}

// UInt64ReadCloser defines an interface for reading a single uint64 type.
type UInt64ReadCloser interface {
	Closer
	UInt64Reader
}

// UInt64StreamReader defines an interface for reading a slice of uint64 type.
type UInt64StreamReader interface {
	Read(int) ([]uint64, error)
}

// UInt64StreamReadCloser defines an interface for reading a single uint64 type.
type UInt64StreamReadCloser interface {
	Closer
	UInt64StreamReader
}

// UInt64Writer defines an interface for writing a single uint64 type.
type UInt64Writer interface {
	WriteUInt64(uint64) (int, error)
}

// UInt64WriteCloser defines an interface for writing a single uint64 type.
type UInt64WriteCloser interface {
	Closer
	UInt64Writer
}

// UInt64StreamWriter defines an interface for writing a slice of uint64 type.
type UInt64StreamWriter interface {
	Write([]uint64) (int, error)
}

// UInt64StreamWriteCloser defines an interface for writing a single uint64 type.
type UInt64StreamWriteCloser interface {
	Closer
	UInt64StreamWriter
}
