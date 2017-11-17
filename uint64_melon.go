package melon

import (
	"github.com/influx6/faux/context"
)

// UInt64UnitReaderFunc defines a function which expects the giving UInt64Reader type has input.
type UInt64UnitReaderFunc func(UInt64Reader) UInt64UnitReader

// UInt64UnitReaderFuncWithContext defines a function which expects the giving UInt64Reader type has input.
// This expects to receive a context.Context type.
type UInt64UnitReaderFuncWithContext func(context.Context, UInt64Reader) UInt64UnitReader

// UInt64ReaderFunc defines a function which expects the giving UInt64Reader type has input.
type UInt64ReaderFunc func(UInt64Reader) error

// UInt64ReaderFuncWithContext defines a function which expects the giving UInt64Reader type has input.
// This expects to receive a context.Context type.
type UInt64ReaderFuncWithContext func(context.Context, UInt64Reader) error

// UInt64ReadCloserFunc defines a function which expects the giving UInt64ReadCloser type has input.
type UInt64ReadCloserFunc func(UInt64ReadCloser) error

// UInt64ReadCloserFuncWithContext defines a function which expects the giving UInt64ReadCloser type has input.
// This expects to receive a context.Context type.
type UInt64ReadCloserFuncWithContext func(context.Context, UInt64ReadCloser) error

// UInt64WriterFunc defines a function which expects the giving UInt64Writer type has input.
type UInt64WriterFunc func(UInt64Writer) error

// UInt64WriteCloserFunc defines a function which expects the giving UInt64WriteCloser type has input.
type UInt64WriteCloserFunc func(UInt64WriteCloser) error

// UInt64WriterFuncWithContext defines a function which expects the giving UInt64Writer type has input.
// This expects to receive a context.Context type.
type UInt64WriterFuncWithContext func(context.Context, UInt64Writer) error

// UInt64WriteCloserFuncWithContext defines a function which expects the giving UInt64WriteCloser type has input.
// This expects to receive a context.Context type.
type UInt64WriteCloserFuncWithContext func(context.Context, UInt64WriteCloser) error

// UInt64Reader defines an interface for reading a slice of uint64 types.
type UInt64Reader interface {
	Read([]uint64) (int, error)
}

// UInt64ReadCloser defines an interface for reading a slice of uint64 types.
type UInt64ReadCloser interface {
	Closer
	UInt64Reader
}

// UInt64UnitReader defines an interface for reading a single item of uint64 type.
type UInt64UnitReader interface {
	Read() (uint64, error)
}

// UInt64Writer defines an interface for writing a slice of uint64 types.
type UInt64Writer interface {
	Write([]uint64) (int, error)
}

// UInt64WriteCloser defines an interface for writing a slice of uint64 types.
type UInt64WriteCloser interface {
	Closer
	UInt64Writer
}
