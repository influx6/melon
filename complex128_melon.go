package melon

import (
	"github.com/influx6/faux/context"
)

// Complex128UniqueHash defines a unique hash for Complex128 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Complex128UniqueHash = "1e22bb114616c1ea2c888d34d5d519a35cc9103c"

// Complex128Handler defines a function type receiving both reader and writer types.
type Complex128Handler func(Complex128Reader, Complex128Writer) error

// Complex128HandlerWithCtx defines a function type receiving both reader and writer types.
type Complex128HandlerWithCtx func(context.Context, Complex128Reader, Complex128Writer) error

// Complex128Stream defines a function type receiving both reader and writer types.
type Complex128Stream func(Complex128StreamReader, Complex128StreamWriter) error

// Complex128StreamWithCtx defines a function type receiving both reader and writer types.
type Complex128StreamWithCtx func(context.Context, Complex128StreamReader, Complex128StreamWriter) error

// Complex128Reader defines an interface for reading a single complex128 type.
type Complex128Reader interface {
	ReadComplex128() (complex128, error)
}

// Complex128ReadCloser defines an interface for reading a single complex128 type.
type Complex128ReadCloser interface {
	Closer
	Complex128Reader
}

// Complex128StreamReader defines an interface for reading a slice of complex128 type.
type Complex128StreamReader interface {
	Read(int) ([]complex128, error)
}

// Complex128StreamReadCloser defines an interface for reading a single complex128 type.
type Complex128StreamReadCloser interface {
	Closer
	Complex128StreamReader
}

// Complex128Writer defines an interface for writing a single complex128 type.
type Complex128Writer interface {
	WriteComplex128(complex128) (int, error)
}

// Complex128WriteCloser defines an interface for writing a single complex128 type.
type Complex128WriteCloser interface {
	Closer
	Complex128Writer
}

// Complex128StreamWriter defines an interface for writing a slice of complex128 type.
type Complex128StreamWriter interface {
	Write([]complex128) (int, error)
}

// Complex128StreamWriteCloser defines an interface for writing a single complex128 type.
type Complex128StreamWriteCloser interface {
	Closer
	Complex128StreamWriter
}
