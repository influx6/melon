package melon

import (
	"github.com/influx6/faux/context"
)

// ErrorUniqueHash defines a unique hash for Error which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const ErrorUniqueHash = "d3e8aa638ed54763fe05bd5404e80abdfc12b4f5"

// ErrorUnitReaderFunc defines a function which expects the giving ErrorReader type has input.
type ErrorUnitReaderFunc func(ErrorReader) ErrorUnitReader

// ErrorUnitReaderFuncWithContext defines a function which expects the giving ErrorReader type has input.
// This expects to receive a context.Context type.
type ErrorUnitReaderFuncWithContext func(context.Context, ErrorReader) ErrorUnitReader

// ErrorReaderFunc defines a function which expects the giving ErrorReader type has input.
type ErrorReaderFunc func(ErrorReader) error

// ErrorReaderFuncWithContext defines a function which expects the giving ErrorReader type has input.
// This expects to receive a context.Context type.
type ErrorReaderFuncWithContext func(context.Context, ErrorReader) error

// ErrorReadCloserFunc defines a function which expects the giving ErrorReadCloser type has input.
type ErrorReadCloserFunc func(ErrorReadCloser) error

// ErrorReadCloserFuncWithContext defines a function which expects the giving ErrorReadCloser type has input.
// This expects to receive a context.Context type.
type ErrorReadCloserFuncWithContext func(context.Context, ErrorReadCloser) error

// ErrorWriterFunc defines a function which expects the giving ErrorWriter type has input.
type ErrorWriterFunc func(ErrorWriter) error

// ErrorWriteCloserFunc defines a function which expects the giving ErrorWriteCloser type has input.
type ErrorWriteCloserFunc func(ErrorWriteCloser) error

// ErrorWriterFuncWithContext defines a function which expects the giving ErrorWriter type has input.
// This expects to receive a context.Context type.
type ErrorWriterFuncWithContext func(context.Context, ErrorWriter) error

// ErrorWriteCloserFuncWithContext defines a function which expects the giving ErrorWriteCloser type has input.
// This expects to receive a context.Context type.
type ErrorWriteCloserFuncWithContext func(context.Context, ErrorWriteCloser) error

// ErrorReader defines an interface for reading a slice of error types.
type ErrorReader interface {
	Read([]error) (int, error)
}

// ErrorReadCloser defines an interface for reading a slice of error types.
type ErrorReadCloser interface {
	Closer
	ErrorReader
}

// ErrorUnitReader defines an interface for reading a single item of error type.
type ErrorUnitReader interface {
	ReadUnit() (error, error)
}

// ErrorWriter defines an interface for writing a slice of error types.
type ErrorWriter interface {
	Write([]error) (int, error)
}

// ErrorUnitWriter defines an interface for writing a single error type.
type ErrorUnitWriter interface {
	WriteUnit(error) (int, error)
}

// ErrorWriteCloser defines an interface for writing a slice of error types.
type ErrorWriteCloser interface {
	Closer
	ErrorWriter
}
