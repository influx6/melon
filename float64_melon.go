package melon

import (
	"github.com/influx6/faux/context"
)

// Float64UnitReaderFunc defines a function which expects the giving Float64Reader type has input.
type Float64UnitReaderFunc func(Float64Reader) Float64UnitReader

// Float64UnitReaderFuncWithContext defines a function which expects the giving Float64Reader type has input.
// This expects to receive a context.Context type.
type Float64UnitReaderFuncWithContext func(context.Context, Float64Reader) Float64UnitReader

// Float64ReaderFunc defines a function which expects the giving Float64Reader type has input.
type Float64ReaderFunc func(Float64Reader) error

// Float64ReaderFuncWithContext defines a function which expects the giving Float64Reader type has input.
// This expects to receive a context.Context type.
type Float64ReaderFuncWithContext func(context.Context, Float64Reader) error

// Float64ReadCloserFunc defines a function which expects the giving Float64ReadCloser type has input.
type Float64ReadCloserFunc func(Float64ReadCloser) error

// Float64ReadCloserFuncWithContext defines a function which expects the giving Float64ReadCloser type has input.
// This expects to receive a context.Context type.
type Float64ReadCloserFuncWithContext func(context.Context, Float64ReadCloser) error

// Float64WriterFunc defines a function which expects the giving Float64Writer type has input.
type Float64WriterFunc func(Float64Writer) error

// Float64WriteCloserFunc defines a function which expects the giving Float64WriteCloser type has input.
type Float64WriteCloserFunc func(Float64WriteCloser) error

// Float64WriterFuncWithContext defines a function which expects the giving Float64Writer type has input.
// This expects to receive a context.Context type.
type Float64WriterFuncWithContext func(context.Context, Float64Writer) error

// Float64WriteCloserFuncWithContext defines a function which expects the giving Float64WriteCloser type has input.
// This expects to receive a context.Context type.
type Float64WriteCloserFuncWithContext func(context.Context, Float64WriteCloser) error

// Float64Reader defines an interface for reading a slice of float64 types.
type Float64Reader interface {
	Read([]float64) (int, error)
}

// Float64ReadCloser defines an interface for reading a slice of float64 types.
type Float64ReadCloser interface {
	Closer
	Float64Reader
}

// Float64UnitReader defines an interface for reading a single item of float64 type.
type Float64UnitReader interface {
	ReadUnit() (float64, error)
}

// Float64Writer defines an interface for writing a slice of float64 types.
type Float64Writer interface {
	Write([]float64) (int, error)
}

// Float64UnitWriter defines an interface for writing a single float64 type.
type Float64UnitWriter interface {
	WriteUnit(float64) (int, error)
}

// Float64WriteCloser defines an interface for writing a slice of float64 types.
type Float64WriteCloser interface {
	Closer
	Float64Writer
}
