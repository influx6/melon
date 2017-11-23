package melon

import (
	"github.com/influx6/faux/context"
)

// MapOfStringUniqueHash defines a unique hash for MapOfString which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const MapOfStringUniqueHash = "de4cf6eba7a008b25c4561bf3c1885a2c02fc96f"

// MapOfStringHandler defines a function type receiving both reader and writer types.
type MapOfStringHandler func(MapOfStringReader, MapOfStringWriter) error

// MapOfStringHandlerWithCtx defines a function type receiving both reader and writer types.
type MapOfStringHandlerWithCtx func(context.Context, MapOfStringReader, MapOfStringWriter) error

// MapOfStringStream defines a function type receiving both reader and writer types.
type MapOfStringStream func(MapOfStringStreamReader, MapOfStringStreamWriter) error

// MapOfStringStreamWithCtx defines a function type receiving both reader and writer types.
type MapOfStringStreamWithCtx func(context.Context, MapOfStringStreamReader, MapOfStringStreamWriter) error

// MapOfStringReader defines an interface for reading a single map[string]string type.
type MapOfStringReader interface {
	ReadMapOfString() (map[string]string, error)
}

// MapOfStringReadCloser defines an interface for reading a single map[string]string type.
type MapOfStringReadCloser interface {
	Closer
	MapOfStringReader
}

// MapOfStringStreamReader defines an interface for reading a slice of map[string]string type.
type MapOfStringStreamReader interface {
	Read(int) ([]map[string]string, error)
}

// MapOfStringStreamReadCloser defines an interface for reading a single map[string]string type.
type MapOfStringStreamReadCloser interface {
	Closer
	MapOfStringStreamReader
}

// MapOfStringWriter defines an interface for writing a single map[string]string type.
type MapOfStringWriter interface {
	WriteMapOfString(map[string]string) (int, error)
}

// MapOfStringWriteCloser defines an interface for writing a single map[string]string type.
type MapOfStringWriteCloser interface {
	Closer
	MapOfStringWriter
}

// MapOfStringStreamWriter defines an interface for writing a slice of map[string]string type.
type MapOfStringStreamWriter interface {
	Write([]map[string]string) (int, error)
}

// MapOfStringStreamWriteCloser defines an interface for writing a single map[string]string type.
type MapOfStringStreamWriteCloser interface {
	Closer
	MapOfStringStreamWriter
}
