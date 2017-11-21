package melon

// Int64UniqueHash defines a unique hash for Int64 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Int64UniqueHash = "d9735f419182b6c737eba2e4b2d3f05a5b349bc9"

// Int64Reader defines an interface for reading a single int64 type.
type Int64Reader interface {
	Read() (int64, error)
}

// Int64LimitReader defines an interface for reading a slice of int64 type.
type Int64LimitReader interface {
	ReadLimit(int) ([]int64, error)
}

// Int64ReadCloser defines an interface for reading a single int64 type.
type Int64ReadCloser interface {
	Closer
	Int64Reader
}

// Int64Writer defines an interface for writing a single int64 type.
type Int64Writer interface {
	Write(int64) (int, error)
}

// Int64LimitWriter defines an interface for writing a slice of int64 type.
type Int64LimitWriter interface {
	WriteLimit([]int64) (int, error)
}

// Int64WriteCloser defines an interface for writing a single int64 type.
type Int64WriteCloser interface {
	Closer
	Int64Writer
}
