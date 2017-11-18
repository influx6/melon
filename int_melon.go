package melon

import (
	"github.com/influx6/faux/context"
)

// IntUniqueHash defines a unique hash for Int which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const IntUniqueHash = "e3e6a2908db2b977524568702fabd849a5d5514f"

// IntUnitReaderFunc defines a function which expects the giving IntReader type has input.
type IntUnitReaderFunc func(IntReader) IntUnitReader

// IntUnitReaderFuncWithContext defines a function which expects the giving IntReader type has input.
// This expects to receive a context.Context type.
type IntUnitReaderFuncWithContext func(context.Context, IntReader) IntUnitReader

// IntReaderFunc defines a function which expects the giving IntReader type has input.
type IntReaderFunc func(IntReader) error

// IntReaderFuncWithContext defines a function which expects the giving IntReader type has input.
// This expects to receive a context.Context type.
type IntReaderFuncWithContext func(context.Context, IntReader) error

// IntReadCloserFunc defines a function which expects the giving IntReadCloser type has input.
type IntReadCloserFunc func(IntReadCloser) error

// IntReadCloserFuncWithContext defines a function which expects the giving IntReadCloser type has input.
// This expects to receive a context.Context type.
type IntReadCloserFuncWithContext func(context.Context, IntReadCloser) error

// IntWriterFunc defines a function which expects the giving IntWriter type has input.
type IntWriterFunc func(IntWriter) error

// IntWriteCloserFunc defines a function which expects the giving IntWriteCloser type has input.
type IntWriteCloserFunc func(IntWriteCloser) error

// IntWriterFuncWithContext defines a function which expects the giving IntWriter type has input.
// This expects to receive a context.Context type.
type IntWriterFuncWithContext func(context.Context, IntWriter) error

// IntWriteCloserFuncWithContext defines a function which expects the giving IntWriteCloser type has input.
// This expects to receive a context.Context type.
type IntWriteCloserFuncWithContext func(context.Context, IntWriteCloser) error

// IntReader defines an interface for reading a slice of int types.
type IntReader interface {
	Read([]int) (int, error)
}

// IntReadCloser defines an interface for reading a slice of int types.
type IntReadCloser interface {
	Closer
	IntReader
}

// IntUnitReader defines an interface for reading a single item of int type.
type IntUnitReader interface {
	ReadUnit() (int, error)
}

// IntWriter defines an interface for writing a slice of int types.
type IntWriter interface {
	Write([]int) (int, error)
}

// IntUnitWriter defines an interface for writing a single int type.
type IntUnitWriter interface {
	WriteUnit(int) (int, error)
}

// IntWriteCloser defines an interface for writing a slice of int types.
type IntWriteCloser interface {
	Closer
	IntWriter
}
