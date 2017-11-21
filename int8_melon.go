package melon

// Int8UniqueHash defines a unique hash for Int8 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Int8UniqueHash = "da945acc637aec1ea9d814cc20b2ed6e794ef8b1"

// Int8Reader defines an interface for reading a single int8 type.
type Int8Reader interface {
	Read() (int8, error)
}

// Int8LimitReader defines an interface for reading a slice of int8 type.
type Int8LimitReader interface {
	ReadLimit(int) ([]int8, error)
}

// Int8ReadCloser defines an interface for reading a single int8 type.
type Int8ReadCloser interface {
	Closer
	Int8Reader
}

// Int8Writer defines an interface for writing a single int8 type.
type Int8Writer interface {
	Write(int8) (int, error)
}

// Int8LimitWriter defines an interface for writing a slice of int8 type.
type Int8LimitWriter interface {
	WriteLimit([]int8) (int, error)
}

// Int8WriteCloser defines an interface for writing a single int8 type.
type Int8WriteCloser interface {
	Closer
	Int8Writer
}
