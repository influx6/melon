package melon

import (
	"github.com/influx6/faux/context"
)

// Int64UniqueHash defines a unique hash for Int64 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Int64UniqueHash = "d9735f419182b6c737eba2e4b2d3f05a5b349bc9"

// Int64UnitReaderFunc defines a function which expects the giving Int64Reader type has input.
type Int64UnitReaderFunc func(Int64Reader) Int64UnitReader

// Int64UnitReaderFuncWithContext defines a function which expects the giving Int64Reader type has input.
// This expects to receive a context.Context type.
type Int64UnitReaderFuncWithContext func(context.Context, Int64Reader) Int64UnitReader

// Int64ReaderFunc defines a function which expects the giving Int64Reader type has input.
type Int64ReaderFunc func(Int64Reader) error

// Int64ReaderFuncWithContext defines a function which expects the giving Int64Reader type has input.
// This expects to receive a context.Context type.
type Int64ReaderFuncWithContext func(context.Context, Int64Reader) error

// Int64ReadCloserFunc defines a function which expects the giving Int64ReadCloser type has input.
type Int64ReadCloserFunc func(Int64ReadCloser) error

// Int64ReadCloserFuncWithContext defines a function which expects the giving Int64ReadCloser type has input.
// This expects to receive a context.Context type.
type Int64ReadCloserFuncWithContext func(context.Context, Int64ReadCloser) error

// Int64WriterFunc defines a function which expects the giving Int64Writer type has input.
type Int64WriterFunc func(Int64Writer) error

// Int64WriteCloserFunc defines a function which expects the giving Int64WriteCloser type has input.
type Int64WriteCloserFunc func(Int64WriteCloser) error

// Int64WriterFuncWithContext defines a function which expects the giving Int64Writer type has input.
// This expects to receive a context.Context type.
type Int64WriterFuncWithContext func(context.Context, Int64Writer) error

// Int64WriteCloserFuncWithContext defines a function which expects the giving Int64WriteCloser type has input.
// This expects to receive a context.Context type.
type Int64WriteCloserFuncWithContext func(context.Context, Int64WriteCloser) error

// Int64Reader defines an interface for reading a slice of int64 types.
type Int64Reader interface {
	Read([]int64) (int, error)
}

// Int64ReadCloser defines an interface for reading a slice of int64 types.
type Int64ReadCloser interface {
	Closer
	Int64Reader
}

// Int64UnitReader defines an interface for reading a single item of int64 type.
type Int64UnitReader interface {
	ReadUnit() (int64, error)
}

// Int64Writer defines an interface for writing a slice of int64 types.
type Int64Writer interface {
	Write([]int64) (int, error)
}

// Int64UnitWriter defines an interface for writing a single int64 type.
type Int64UnitWriter interface {
	WriteUnit(int64) (int, error)
}

// Int64WriteCloser defines an interface for writing a slice of int64 types.
type Int64WriteCloser interface {
	Closer
	Int64Writer
}
