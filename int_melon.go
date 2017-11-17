package melon

import (
	"github.com/influx6/faux/context"
)

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

// IntWriter defines an interface for writing a slice of int types.
type IntWriter interface {
	Write([]int) (int, error)
}

// IntWriteCloser defines an interface for writing a slice of int types.
type IntWriteCloser interface {
	Closer
	IntWriter
}
