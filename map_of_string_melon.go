package melon

import (
	"github.com/influx6/faux/context"
)

// MapOfStringReaderFunc defines a function which expects the giving MapOfStringReader type has input.
type MapOfStringReaderFunc func(MapOfStringReader) error

// MapOfStringReaderFuncWithContext defines a function which expects the giving MapOfStringReader type has input.
// This expects to receive a context.Context type.
type MapOfStringReaderFuncWithContext func(context.Context, MapOfStringReader) error

// MapOfStringReadCloserFunc defines a function which expects the giving MapOfStringReadCloser type has input.
type MapOfStringReadCloserFunc func(MapOfStringReadCloser) error

// MapOfStringReadCloserFuncWithContext defines a function which expects the giving MapOfStringReadCloser type has input.
// This expects to receive a context.Context type.
type MapOfStringReadCloserFuncWithContext func(context.Context, MapOfStringReadCloser) error

// MapOfStringWriterFunc defines a function which expects the giving MapOfStringWriter type has input.
type MapOfStringWriterFunc func(MapOfStringWriter) error

// MapOfStringWriteCloserFunc defines a function which expects the giving MapOfStringWriteCloser type has input.
type MapOfStringWriteCloserFunc func(MapOfStringWriteCloser) error

// MapOfStringWriterFuncWithContext defines a function which expects the giving MapOfStringWriter type has input.
// This expects to receive a context.Context type.
type MapOfStringWriterFuncWithContext func(context.Context, MapOfStringWriter) error

// MapOfStringWriteCloserFuncWithContext defines a function which expects the giving MapOfStringWriteCloser type has input.
// This expects to receive a context.Context type.
type MapOfStringWriteCloserFuncWithContext func(context.Context, MapOfStringWriteCloser) error

// MapOfStringReader defines an interface for reading a slice of map[string]interface{} types.
type MapOfStringReader interface {
	Read([]map[string]interface{}) (int, error)
}

// MapOfStringReadCloser defines an interface for reading a slice of map[string]interface{} types.
type MapOfStringReadCloser interface {
	Closer
	MapOfStringReader
}

// MapOfStringWriter defines an interface for writing a slice of map[string]interface{} types.
type MapOfStringWriter interface {
	Write([]map[string]interface{}) (int, error)
}

// MapOfStringWriteCloser defines an interface for writing a slice of map[string]interface{} types.
type MapOfStringWriteCloser interface {
	Closer
	MapOfStringWriter
}
