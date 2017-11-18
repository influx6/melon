package melon

import (
	"github.com/influx6/faux/context"
)

// MapOfStringUniqueHash defines a unique hash for MapOfString which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const MapOfStringUniqueHash = "de4cf6eba7a008b25c4561bf3c1885a2c02fc96f"

// MapOfStringUnitReaderFunc defines a function which expects the giving MapOfStringReader type has input.
type MapOfStringUnitReaderFunc func(MapOfStringReader) MapOfStringUnitReader

// MapOfStringUnitReaderFuncWithContext defines a function which expects the giving MapOfStringReader type has input.
// This expects to receive a context.Context type.
type MapOfStringUnitReaderFuncWithContext func(context.Context, MapOfStringReader) MapOfStringUnitReader

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

// MapOfStringReader defines an interface for reading a slice of map[string]string types.
type MapOfStringReader interface {
	Read([]map[string]string) (int, error)
}

// MapOfStringReadCloser defines an interface for reading a slice of map[string]string types.
type MapOfStringReadCloser interface {
	Closer
	MapOfStringReader
}

// MapOfStringUnitReader defines an interface for reading a single item of map[string]string type.
type MapOfStringUnitReader interface {
	ReadUnit() (map[string]string, error)
}

// MapOfStringWriter defines an interface for writing a slice of map[string]string types.
type MapOfStringWriter interface {
	Write([]map[string]string) (int, error)
}

// MapOfStringUnitWriter defines an interface for writing a single map[string]string type.
type MapOfStringUnitWriter interface {
	WriteUnit(map[string]string) (int, error)
}

// MapOfStringWriteCloser defines an interface for writing a slice of map[string]string types.
type MapOfStringWriteCloser interface {
	Closer
	MapOfStringWriter
}
