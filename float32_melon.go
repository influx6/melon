package melon

import (
	"github.com/influx6/faux/context"
)

// Float32UnitReaderFunc defines a function which expects the giving Float32Reader type has input.
type Float32UnitReaderFunc func(Float32Reader) Float32UnitReader

// Float32UnitReaderFuncWithContext defines a function which expects the giving Float32Reader type has input.
// This expects to receive a context.Context type.
type Float32UnitReaderFuncWithContext func(context.Context, Float32Reader) Float32UnitReader

// Float32ReaderFunc defines a function which expects the giving Float32Reader type has input.
type Float32ReaderFunc func(Float32Reader) error

// Float32ReaderFuncWithContext defines a function which expects the giving Float32Reader type has input.
// This expects to receive a context.Context type.
type Float32ReaderFuncWithContext func(context.Context, Float32Reader) error

// Float32ReadCloserFunc defines a function which expects the giving Float32ReadCloser type has input.
type Float32ReadCloserFunc func(Float32ReadCloser) error

// Float32ReadCloserFuncWithContext defines a function which expects the giving Float32ReadCloser type has input.
// This expects to receive a context.Context type.
type Float32ReadCloserFuncWithContext func(context.Context, Float32ReadCloser) error

// Float32WriterFunc defines a function which expects the giving Float32Writer type has input.
type Float32WriterFunc func(Float32Writer) error

// Float32WriteCloserFunc defines a function which expects the giving Float32WriteCloser type has input.
type Float32WriteCloserFunc func(Float32WriteCloser) error

// Float32WriterFuncWithContext defines a function which expects the giving Float32Writer type has input.
// This expects to receive a context.Context type.
type Float32WriterFuncWithContext func(context.Context, Float32Writer) error

// Float32WriteCloserFuncWithContext defines a function which expects the giving Float32WriteCloser type has input.
// This expects to receive a context.Context type.
type Float32WriteCloserFuncWithContext func(context.Context, Float32WriteCloser) error

// Float32Reader defines an interface for reading a slice of float32 types.
type Float32Reader interface {
	Read([]float32) (int, error)
}

// Float32ReadCloser defines an interface for reading a slice of float32 types.
type Float32ReadCloser interface {
	Closer
	Float32Reader
}

// Float32UnitReader defines an interface for reading a single item of float32 type.
type Float32UnitReader interface {
	Read() (float32, error)
}

// Float32Writer defines an interface for writing a slice of float32 types.
type Float32Writer interface {
	Write([]float32) (int, error)
}

// Float32WriteCloser defines an interface for writing a slice of float32 types.
type Float32WriteCloser interface {
	Closer
	Float32Writer
}
