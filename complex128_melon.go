package melon

import (
	"github.com/influx6/faux/context"
)

// Complex128UniqueHash defines a unique hash for Complex128 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Complex128UniqueHash = "1e22bb114616c1ea2c888d34d5d519a35cc9103c"

// Complex128UnitReaderFunc defines a function which expects the giving Complex128Reader type has input.
type Complex128UnitReaderFunc func(Complex128Reader) Complex128UnitReader

// Complex128UnitReaderFuncWithContext defines a function which expects the giving Complex128Reader type has input.
// This expects to receive a context.Context type.
type Complex128UnitReaderFuncWithContext func(context.Context, Complex128Reader) Complex128UnitReader

// Complex128ReaderFunc defines a function which expects the giving Complex128Reader type has input.
type Complex128ReaderFunc func(Complex128Reader) error

// Complex128ReaderFuncWithContext defines a function which expects the giving Complex128Reader type has input.
// This expects to receive a context.Context type.
type Complex128ReaderFuncWithContext func(context.Context, Complex128Reader) error

// Complex128ReadCloserFunc defines a function which expects the giving Complex128ReadCloser type has input.
type Complex128ReadCloserFunc func(Complex128ReadCloser) error

// Complex128ReadCloserFuncWithContext defines a function which expects the giving Complex128ReadCloser type has input.
// This expects to receive a context.Context type.
type Complex128ReadCloserFuncWithContext func(context.Context, Complex128ReadCloser) error

// Complex128WriterFunc defines a function which expects the giving Complex128Writer type has input.
type Complex128WriterFunc func(Complex128Writer) error

// Complex128WriteCloserFunc defines a function which expects the giving Complex128WriteCloser type has input.
type Complex128WriteCloserFunc func(Complex128WriteCloser) error

// Complex128WriterFuncWithContext defines a function which expects the giving Complex128Writer type has input.
// This expects to receive a context.Context type.
type Complex128WriterFuncWithContext func(context.Context, Complex128Writer) error

// Complex128WriteCloserFuncWithContext defines a function which expects the giving Complex128WriteCloser type has input.
// This expects to receive a context.Context type.
type Complex128WriteCloserFuncWithContext func(context.Context, Complex128WriteCloser) error

// Complex128Reader defines an interface for reading a slice of complex128 types.
type Complex128Reader interface {
	Read([]complex128) (int, error)
}

// Complex128ReadCloser defines an interface for reading a slice of complex128 types.
type Complex128ReadCloser interface {
	Closer
	Complex128Reader
}

// Complex128UnitReader defines an interface for reading a single item of complex128 type.
type Complex128UnitReader interface {
	ReadUnit() (complex128, error)
}

// Complex128Writer defines an interface for writing a slice of complex128 types.
type Complex128Writer interface {
	Write([]complex128) (int, error)
}

// Complex128UnitWriter defines an interface for writing a single complex128 type.
type Complex128UnitWriter interface {
	WriteUnit(complex128) (int, error)
}

// Complex128WriteCloser defines an interface for writing a slice of complex128 types.
type Complex128WriteCloser interface {
	Closer
	Complex128Writer
}
