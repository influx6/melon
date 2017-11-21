package melon

// MapOfAnyUniqueHash defines a unique hash for MapOfAny which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const MapOfAnyUniqueHash = "c2d03452b8860af6906e2c8a6e847f371db07988"

// MapOfAnyReader defines an interface for reading a single map[interface{}]interface{} type.
type MapOfAnyReader interface {
	Read() (map[interface{}]interface{}, error)
}

// MapOfAnyLimitReader defines an interface for reading a slice of map[interface{}]interface{} type.
type MapOfAnyLimitReader interface {
	ReadLimit(int) ([]map[interface{}]interface{}, error)
}

// MapOfAnyReadCloser defines an interface for reading a single map[interface{}]interface{} type.
type MapOfAnyReadCloser interface {
	Closer
	MapOfAnyReader
}

// MapOfAnyWriter defines an interface for writing a single map[interface{}]interface{} type.
type MapOfAnyWriter interface {
	Write(map[interface{}]interface{}) (int, error)
}

// MapOfAnyLimitWriter defines an interface for writing a slice of map[interface{}]interface{} type.
type MapOfAnyLimitWriter interface {
	WriteLimit([]map[interface{}]interface{}) (int, error)
}

// MapOfAnyWriteCloser defines an interface for writing a single map[interface{}]interface{} type.
type MapOfAnyWriteCloser interface {
	Closer
	MapOfAnyWriter
}
