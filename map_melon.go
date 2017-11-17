package melon

import (
	"github.com/influx6/faux/context"
)

// MapUnitReaderFunc defines a function which expects the giving MapReader type has input.
type MapUnitReaderFunc func(MapReader) MapUnitReader

// MapUnitReaderFuncWithContext defines a function which expects the giving MapReader type has input.
// This expects to receive a context.Context type.
type MapUnitReaderFuncWithContext func(context.Context, MapReader) MapUnitReader

// MapReaderFunc defines a function which expects the giving MapReader type has input.
type MapReaderFunc func(MapReader) error

// MapReaderFuncWithContext defines a function which expects the giving MapReader type has input.
// This expects to receive a context.Context type.
type MapReaderFuncWithContext func(context.Context, MapReader) error

// MapReadCloserFunc defines a function which expects the giving MapReadCloser type has input.
type MapReadCloserFunc func(MapReadCloser) error

// MapReadCloserFuncWithContext defines a function which expects the giving MapReadCloser type has input.
// This expects to receive a context.Context type.
type MapReadCloserFuncWithContext func(context.Context, MapReadCloser) error

// MapWriterFunc defines a function which expects the giving MapWriter type has input.
type MapWriterFunc func(MapWriter) error

// MapWriteCloserFunc defines a function which expects the giving MapWriteCloser type has input.
type MapWriteCloserFunc func(MapWriteCloser) error

// MapWriterFuncWithContext defines a function which expects the giving MapWriter type has input.
// This expects to receive a context.Context type.
type MapWriterFuncWithContext func(context.Context, MapWriter) error

// MapWriteCloserFuncWithContext defines a function which expects the giving MapWriteCloser type has input.
// This expects to receive a context.Context type.
type MapWriteCloserFuncWithContext func(context.Context, MapWriteCloser) error

// MapReader defines an interface for reading a slice of map[string]interface{} types.
type MapReader interface {
	Read([]map[string]interface{}) (int, error)
}

// MapReadCloser defines an interface for reading a slice of map[string]interface{} types.
type MapReadCloser interface {
	Closer
	MapReader
}

// MapUnitReader defines an interface for reading a single item of map[string]interface{} type.
type MapUnitReader interface {
	ReadUnit() (map[string]interface{}, error)
}

// MapWriter defines an interface for writing a slice of map[string]interface{} types.
type MapWriter interface {
	Write([]map[string]interface{}) (int, error)
}

// MapUnitWriter defines an interface for writing a single map[string]interface{} type.
type MapUnitWriter interface {
	WriteUnit(map[string]interface{}) (int, error)
}

// MapWriteCloser defines an interface for writing a slice of map[string]interface{} types.
type MapWriteCloser interface {
	Closer
	MapWriter
}
