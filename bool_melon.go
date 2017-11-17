package melon

import (
	"github.com/influx6/faux/context"
)

// BoolUnitReaderFunc defines a function which expects the giving BoolReader type has input.
type BoolUnitReaderFunc func(BoolReader) BoolUnitReader

// BoolUnitReaderFuncWithContext defines a function which expects the giving BoolReader type has input.
// This expects to receive a context.Context type.
type BoolUnitReaderFuncWithContext func(context.Context, BoolReader) BoolUnitReader

// BoolReaderFunc defines a function which expects the giving BoolReader type has input.
type BoolReaderFunc func(BoolReader) error

// BoolReaderFuncWithContext defines a function which expects the giving BoolReader type has input.
// This expects to receive a context.Context type.
type BoolReaderFuncWithContext func(context.Context, BoolReader) error

// BoolReadCloserFunc defines a function which expects the giving BoolReadCloser type has input.
type BoolReadCloserFunc func(BoolReadCloser) error

// BoolReadCloserFuncWithContext defines a function which expects the giving BoolReadCloser type has input.
// This expects to receive a context.Context type.
type BoolReadCloserFuncWithContext func(context.Context, BoolReadCloser) error

// BoolWriterFunc defines a function which expects the giving BoolWriter type has input.
type BoolWriterFunc func(BoolWriter) error

// BoolWriteCloserFunc defines a function which expects the giving BoolWriteCloser type has input.
type BoolWriteCloserFunc func(BoolWriteCloser) error

// BoolWriterFuncWithContext defines a function which expects the giving BoolWriter type has input.
// This expects to receive a context.Context type.
type BoolWriterFuncWithContext func(context.Context, BoolWriter) error

// BoolWriteCloserFuncWithContext defines a function which expects the giving BoolWriteCloser type has input.
// This expects to receive a context.Context type.
type BoolWriteCloserFuncWithContext func(context.Context, BoolWriteCloser) error

// BoolReader defines an interface for reading a slice of bool types.
type BoolReader interface {
	Read([]bool) (int, error)
}

// BoolReadCloser defines an interface for reading a slice of bool types.
type BoolReadCloser interface {
	Closer
	BoolReader
}

// BoolUnitReader defines an interface for reading a single item of bool type.
type BoolUnitReader interface {
	Read() (bool, error)
}

// BoolWriter defines an interface for writing a slice of bool types.
type BoolWriter interface {
	Write([]bool) (int, error)
}

// BoolWriteCloser defines an interface for writing a slice of bool types.
type BoolWriteCloser interface {
	Closer
	BoolWriter
}
