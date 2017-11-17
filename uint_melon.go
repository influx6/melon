package melon

import (
	"github.com/influx6/faux/context"
)

// UIntUnitReaderFunc defines a function which expects the giving UIntReader type has input.
type UIntUnitReaderFunc func(UIntReader) UIntUnitReader

// UIntUnitReaderFuncWithContext defines a function which expects the giving UIntReader type has input.
// This expects to receive a context.Context type.
type UIntUnitReaderFuncWithContext func(context.Context, UIntReader) UIntUnitReader

// UIntReaderFunc defines a function which expects the giving UIntReader type has input.
type UIntReaderFunc func(UIntReader) error

// UIntReaderFuncWithContext defines a function which expects the giving UIntReader type has input.
// This expects to receive a context.Context type.
type UIntReaderFuncWithContext func(context.Context, UIntReader) error

// UIntReadCloserFunc defines a function which expects the giving UIntReadCloser type has input.
type UIntReadCloserFunc func(UIntReadCloser) error

// UIntReadCloserFuncWithContext defines a function which expects the giving UIntReadCloser type has input.
// This expects to receive a context.Context type.
type UIntReadCloserFuncWithContext func(context.Context, UIntReadCloser) error

// UIntWriterFunc defines a function which expects the giving UIntWriter type has input.
type UIntWriterFunc func(UIntWriter) error

// UIntWriteCloserFunc defines a function which expects the giving UIntWriteCloser type has input.
type UIntWriteCloserFunc func(UIntWriteCloser) error

// UIntWriterFuncWithContext defines a function which expects the giving UIntWriter type has input.
// This expects to receive a context.Context type.
type UIntWriterFuncWithContext func(context.Context, UIntWriter) error

// UIntWriteCloserFuncWithContext defines a function which expects the giving UIntWriteCloser type has input.
// This expects to receive a context.Context type.
type UIntWriteCloserFuncWithContext func(context.Context, UIntWriteCloser) error

// UIntReader defines an interface for reading a slice of uint types.
type UIntReader interface {
	Read([]uint) (int, error)
}

// UIntReadCloser defines an interface for reading a slice of uint types.
type UIntReadCloser interface {
	Closer
	UIntReader
}

// UIntUnitReader defines an interface for reading a single item of uint type.
type UIntUnitReader interface {
	Read() (uint, error)
}

// UIntWriter defines an interface for writing a slice of uint types.
type UIntWriter interface {
	Write([]uint) (int, error)
}

// UIntWriteCloser defines an interface for writing a slice of uint types.
type UIntWriteCloser interface {
	Closer
	UIntWriter
}
