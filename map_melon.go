package melon

import (
	"github.com/influx6/faux/context"
)

// MapUniqueHash defines a unique hash for Map which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const MapUniqueHash = "ba70ac0926874b71d829907f558c032da235489c"

// MapHandler defines a function type receiving both reader and writer types.
type MapHandler func(MapReader, MapWriter) error

// MapHandlerWithCtx defines a function type receiving both reader and writer types.
type MapHandlerWithCtx func(context.Context, MapReader, MapWriter) error

// MapStream defines a function type receiving both reader and writer types.
type MapStream func(MapStreamReader, MapStreamWriter) error

// MapStreamWithCtx defines a function type receiving both reader and writer types.
type MapStreamWithCtx func(context.Context, MapStreamReader, MapStreamWriter) error

// MapReader defines an interface for reading a single map[string]interface{} type.
type MapReader interface {
	ReadMap() (map[string]interface{}, error)
}

// MapReadCloser defines an interface for reading a single map[string]interface{} type.
type MapReadCloser interface {
	Closer
	MapReader
}

// MapStreamReader defines an interface for reading a slice of map[string]interface{} type.
type MapStreamReader interface {
	Read(int) ([]map[string]interface{}, error)
}

// MapStreamReadCloser defines an interface for reading a single map[string]interface{} type.
type MapStreamReadCloser interface {
	Closer
	MapStreamReader
}

// MapWriter defines an interface for writing a single map[string]interface{} type.
type MapWriter interface {
	WriteMap(map[string]interface{}) (int, error)
}

// MapWriteCloser defines an interface for writing a single map[string]interface{} type.
type MapWriteCloser interface {
	Closer
	MapWriter
}

// MapStreamWriter defines an interface for writing a slice of map[string]interface{} type.
type MapStreamWriter interface {
	Write([]map[string]interface{}) (int, error)
}

// MapStreamWriteCloser defines an interface for writing a single map[string]interface{} type.
type MapStreamWriteCloser interface {
	Closer
	MapStreamWriter
}
