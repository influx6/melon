package melon

import (
	"github.com/influx6/faux/context"
)

// MapOfAnyUniqueHash defines a unique hash for MapOfAny which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const MapOfAnyUniqueHash = "c2d03452b8860af6906e2c8a6e847f371db07988"

// MapOfAnyHandler defines a function type receiving both reader and writer types.
type MapOfAnyHandler func(MapOfAnyReader, MapOfAnyWriter) error

// MapOfAnyHandlerWithCtx defines a function type receiving both reader and writer types.
type MapOfAnyHandlerWithCtx func(context.Context, MapOfAnyReader, MapOfAnyWriter) error

// MapOfAnyStream defines a function type receiving both reader and writer types.
type MapOfAnyStream func(MapOfAnyStreamReader, MapOfAnyStreamWriter) error

// MapOfAnyStreamWithCtx defines a function type receiving both reader and writer types.
type MapOfAnyStreamWithCtx func(context.Context, MapOfAnyStreamReader, MapOfAnyStreamWriter) error

// MapOfAnyReader defines an interface for reading a single map[interface{}]interface{} type.
type MapOfAnyReader interface {
	ReadMapOfAny() (map[interface{}]interface{}, error)
}

// MapOfAnyReadCloser defines an interface for reading a single map[interface{}]interface{} type.
type MapOfAnyReadCloser interface {
	Closer
	MapOfAnyReader
}

// MapOfAnyStreamReader defines an interface for reading a slice of map[interface{}]interface{} type.
type MapOfAnyStreamReader interface {
	Read(int) ([]map[interface{}]interface{}, error)
}

// MapOfAnyStreamReadCloser defines an interface for reading a single map[interface{}]interface{} type.
type MapOfAnyStreamReadCloser interface {
	Closer
	MapOfAnyStreamReader
}

// MapOfAnyWriter defines an interface for writing a single map[interface{}]interface{} type.
type MapOfAnyWriter interface {
	WriteMapOfAny(map[interface{}]interface{}) (int, error)
}

// MapOfAnyWriteCloser defines an interface for writing a single map[interface{}]interface{} type.
type MapOfAnyWriteCloser interface {
	Closer
	MapOfAnyWriter
}

// MapOfAnyStreamWriter defines an interface for writing a slice of map[interface{}]interface{} type.
type MapOfAnyStreamWriter interface {
	Write([]map[interface{}]interface{}) (int, error)
}

// MapOfAnyStreamWriteCloser defines an interface for writing a single map[interface{}]interface{} type.
type MapOfAnyStreamWriteCloser interface {
	Closer
	MapOfAnyStreamWriter
}
