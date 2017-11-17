package melon

import (
	"github.com/influx6/faux/context"
)

// StringUnitReaderFunc defines a function which expects the giving StringReader type has input.
type StringUnitReaderFunc func(StringReader) StringUnitReader

// StringUnitReaderFuncWithContext defines a function which expects the giving StringReader type has input.
// This expects to receive a context.Context type.
type StringUnitReaderFuncWithContext func(context.Context, StringReader) StringUnitReader

// StringReaderFunc defines a function which expects the giving StringReader type has input.
type StringReaderFunc func(StringReader) error

// StringReaderFuncWithContext defines a function which expects the giving StringReader type has input.
// This expects to receive a context.Context type.
type StringReaderFuncWithContext func(context.Context, StringReader) error

// StringReadCloserFunc defines a function which expects the giving StringReadCloser type has input.
type StringReadCloserFunc func(StringReadCloser) error

// StringReadCloserFuncWithContext defines a function which expects the giving StringReadCloser type has input.
// This expects to receive a context.Context type.
type StringReadCloserFuncWithContext func(context.Context, StringReadCloser) error

// StringWriterFunc defines a function which expects the giving StringWriter type has input.
type StringWriterFunc func(StringWriter) error

// StringWriteCloserFunc defines a function which expects the giving StringWriteCloser type has input.
type StringWriteCloserFunc func(StringWriteCloser) error

// StringWriterFuncWithContext defines a function which expects the giving StringWriter type has input.
// This expects to receive a context.Context type.
type StringWriterFuncWithContext func(context.Context, StringWriter) error

// StringWriteCloserFuncWithContext defines a function which expects the giving StringWriteCloser type has input.
// This expects to receive a context.Context type.
type StringWriteCloserFuncWithContext func(context.Context, StringWriteCloser) error

// StringReader defines an interface for reading a slice of string types.
type StringReader interface {
	Read([]string) (int, error)
}

// StringReadCloser defines an interface for reading a slice of string types.
type StringReadCloser interface {
	Closer
	StringReader
}

// StringUnitReader defines an interface for reading a single item of string type.
type StringUnitReader interface {
	Read() (string, error)
}

// StringWriter defines an interface for writing a slice of string types.
type StringWriter interface {
	Write([]string) (int, error)
}

// StringWriteCloser defines an interface for writing a slice of string types.
type StringWriteCloser interface {
	Closer
	StringWriter
}
