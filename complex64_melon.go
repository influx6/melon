package melon

import (
	"github.com/influx6/faux/context"
)

// Complex64UniqueHash defines a unique hash for Complex64 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Complex64UniqueHash = "161e6337d01e5b8bd4e79ac4060fde11263b2d17"

// Complex64Handler defines a function type receiving both reader and writer types.
type Complex64Handler func(Complex64Reader, Complex64Writer) error

// Complex64HandlerWithCtx defines a function type receiving both reader and writer types.
type Complex64HandlerWithCtx func(context.Context, Complex64Reader, Complex64Writer) error

// Complex64Stream defines a function type receiving both reader and writer types.
type Complex64Stream func(Complex64StreamReader, Complex64StreamWriter) error

// Complex64StreamWithCtx defines a function type receiving both reader and writer types.
type Complex64StreamWithCtx func(context.Context, Complex64StreamReader, Complex64StreamWriter) error

// Complex64Reader defines an interface for reading a single complex64 type.
type Complex64Reader interface {
	ReadComplex64() (complex64, error)
}

// Complex64ReadCloser defines an interface for reading a single complex64 type.
type Complex64ReadCloser interface {
	Closer
	Complex64Reader
}

// Complex64StreamReader defines an interface for reading a slice of complex64 type.
type Complex64StreamReader interface {
	Read(int) ([]complex64, error)
}

// Complex64StreamReadCloser defines an interface for reading a single complex64 type.
type Complex64StreamReadCloser interface {
	Closer
	Complex64StreamReader
}

// Complex64Writer defines an interface for writing a single complex64 type.
type Complex64Writer interface {
	WriteComplex64(complex64) (int, error)
}

// Complex64WriteCloser defines an interface for writing a single complex64 type.
type Complex64WriteCloser interface {
	Closer
	Complex64Writer
}

// Complex64StreamWriter defines an interface for writing a slice of complex64 type.
type Complex64StreamWriter interface {
	Write([]complex64) (int, error)
}

// Complex64StreamWriteCloser defines an interface for writing a single complex64 type.
type Complex64StreamWriteCloser interface {
	Closer
	Complex64StreamWriter
}
