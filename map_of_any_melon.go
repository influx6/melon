package melon

import (
	"github.com/influx6/faux/context"
)

// MapOfAnyUnitReaderFunc defines a function which expects the giving MapOfAnyReader type has input.
type MapOfAnyUnitReaderFunc func(MapOfAnyReader) MapOfAnyUnitReader

// MapOfAnyUnitReaderFuncWithContext defines a function which expects the giving MapOfAnyReader type has input.
// This expects to receive a context.Context type.
type MapOfAnyUnitReaderFuncWithContext func(context.Context, MapOfAnyReader) MapOfAnyUnitReader

// MapOfAnyReaderFunc defines a function which expects the giving MapOfAnyReader type has input.
type MapOfAnyReaderFunc func(MapOfAnyReader) error

// MapOfAnyReaderFuncWithContext defines a function which expects the giving MapOfAnyReader type has input.
// This expects to receive a context.Context type.
type MapOfAnyReaderFuncWithContext func(context.Context, MapOfAnyReader) error

// MapOfAnyReadCloserFunc defines a function which expects the giving MapOfAnyReadCloser type has input.
type MapOfAnyReadCloserFunc func(MapOfAnyReadCloser) error

// MapOfAnyReadCloserFuncWithContext defines a function which expects the giving MapOfAnyReadCloser type has input.
// This expects to receive a context.Context type.
type MapOfAnyReadCloserFuncWithContext func(context.Context, MapOfAnyReadCloser) error

// MapOfAnyWriterFunc defines a function which expects the giving MapOfAnyWriter type has input.
type MapOfAnyWriterFunc func(MapOfAnyWriter) error

// MapOfAnyWriteCloserFunc defines a function which expects the giving MapOfAnyWriteCloser type has input.
type MapOfAnyWriteCloserFunc func(MapOfAnyWriteCloser) error

// MapOfAnyWriterFuncWithContext defines a function which expects the giving MapOfAnyWriter type has input.
// This expects to receive a context.Context type.
type MapOfAnyWriterFuncWithContext func(context.Context, MapOfAnyWriter) error

// MapOfAnyWriteCloserFuncWithContext defines a function which expects the giving MapOfAnyWriteCloser type has input.
// This expects to receive a context.Context type.
type MapOfAnyWriteCloserFuncWithContext func(context.Context, MapOfAnyWriteCloser) error

// MapOfAnyReader defines an interface for reading a slice of map[interface{}]interface{} types.
type MapOfAnyReader interface {
	Read([]map[interface{}]interface{}) (int, error)
}

// MapOfAnyReadCloser defines an interface for reading a slice of map[interface{}]interface{} types.
type MapOfAnyReadCloser interface {
	Closer
	MapOfAnyReader
}

// MapOfAnyUnitReader defines an interface for reading a single item of map[interface{}]interface{} type.
type MapOfAnyUnitReader interface {
	Read() (map[interface{}]interface{}, error)
}

// MapOfAnyWriter defines an interface for writing a slice of map[interface{}]interface{} types.
type MapOfAnyWriter interface {
	Write([]map[interface{}]interface{}) (int, error)
}

// MapOfAnyWriteCloser defines an interface for writing a slice of map[interface{}]interface{} types.
type MapOfAnyWriteCloser interface {
	Closer
	MapOfAnyWriter
}
