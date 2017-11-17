package melon

import (
	"github.com/influx6/faux/context"
)

// Complex64UnitReaderFunc defines a function which expects the giving Complex64Reader type has input.
type Complex64UnitReaderFunc func(Complex64Reader) Complex64UnitReader

// Complex64UnitReaderFuncWithContext defines a function which expects the giving Complex64Reader type has input.
// This expects to receive a context.Context type.
type Complex64UnitReaderFuncWithContext func(context.Context, Complex64Reader) Complex64UnitReader

// Complex64ReaderFunc defines a function which expects the giving Complex64Reader type has input.
type Complex64ReaderFunc func(Complex64Reader) error

// Complex64ReaderFuncWithContext defines a function which expects the giving Complex64Reader type has input.
// This expects to receive a context.Context type.
type Complex64ReaderFuncWithContext func(context.Context, Complex64Reader) error

// Complex64ReadCloserFunc defines a function which expects the giving Complex64ReadCloser type has input.
type Complex64ReadCloserFunc func(Complex64ReadCloser) error

// Complex64ReadCloserFuncWithContext defines a function which expects the giving Complex64ReadCloser type has input.
// This expects to receive a context.Context type.
type Complex64ReadCloserFuncWithContext func(context.Context, Complex64ReadCloser) error

// Complex64WriterFunc defines a function which expects the giving Complex64Writer type has input.
type Complex64WriterFunc func(Complex64Writer) error

// Complex64WriteCloserFunc defines a function which expects the giving Complex64WriteCloser type has input.
type Complex64WriteCloserFunc func(Complex64WriteCloser) error

// Complex64WriterFuncWithContext defines a function which expects the giving Complex64Writer type has input.
// This expects to receive a context.Context type.
type Complex64WriterFuncWithContext func(context.Context, Complex64Writer) error

// Complex64WriteCloserFuncWithContext defines a function which expects the giving Complex64WriteCloser type has input.
// This expects to receive a context.Context type.
type Complex64WriteCloserFuncWithContext func(context.Context, Complex64WriteCloser) error

// Complex64Reader defines an interface for reading a slice of complex64 types.
type Complex64Reader interface {
	Read([]complex64) (int, error)
}

// Complex64ReadCloser defines an interface for reading a slice of complex64 types.
type Complex64ReadCloser interface {
	Closer
	Complex64Reader
}

// Complex64UnitReader defines an interface for reading a single item of complex64 type.
type Complex64UnitReader interface {
	Read() (complex64, error)
}

// Complex64Writer defines an interface for writing a slice of complex64 types.
type Complex64Writer interface {
	Write([]complex64) (int, error)
}

// Complex64WriteCloser defines an interface for writing a slice of complex64 types.
type Complex64WriteCloser interface {
	Closer
	Complex64Writer
}
