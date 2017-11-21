package melon

// MapUniqueHash defines a unique hash for Map which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const MapUniqueHash = "ba70ac0926874b71d829907f558c032da235489c"

// MapReader defines an interface for reading a single map[string]interface{} type.
type MapReader interface {
	Read() (map[string]interface{}, error)
}

// MapLimitReader defines an interface for reading a slice of map[string]interface{} type.
type MapLimitReader interface {
	ReadLimit(int) ([]map[string]interface{}, error)
}

// MapReadCloser defines an interface for reading a single map[string]interface{} type.
type MapReadCloser interface {
	Closer
	MapReader
}

// MapWriter defines an interface for writing a single map[string]interface{} type.
type MapWriter interface {
	Write(map[string]interface{}) (int, error)
}

// MapLimitWriter defines an interface for writing a slice of map[string]interface{} type.
type MapLimitWriter interface {
	WriteLimit([]map[string]interface{}) (int, error)
}

// MapWriteCloser defines an interface for writing a single map[string]interface{} type.
type MapWriteCloser interface {
	Closer
	MapWriter
}
