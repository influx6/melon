package melon

// Int16UniqueHash defines a unique hash for Int16 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Int16UniqueHash = "3b1227b09fd0719b0d9a553739bfa14479fc0a09"

// Int16Reader defines an interface for reading a single int16 type.
type Int16Reader interface {
	Read() (int16, error)
}

// Int16LimitReader defines an interface for reading a slice of int16 type.
type Int16LimitReader interface {
	ReadLimit(int) ([]int16, error)
}

// Int16ReadCloser defines an interface for reading a single int16 type.
type Int16ReadCloser interface {
	Closer
	Int16Reader
}

// Int16Writer defines an interface for writing a single int16 type.
type Int16Writer interface {
	Write(int16) (int, error)
}

// Int16LimitWriter defines an interface for writing a slice of int16 type.
type Int16LimitWriter interface {
	WriteLimit([]int16) (int, error)
}

// Int16WriteCloser defines an interface for writing a single int16 type.
type Int16WriteCloser interface {
	Closer
	Int16Writer
}
