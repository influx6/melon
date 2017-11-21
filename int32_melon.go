package melon

// Int32UniqueHash defines a unique hash for Int32 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Int32UniqueHash = "f90a5efe8610965bee01abba15d0e1f54046e2e8"

// Int32Reader defines an interface for reading a single int32 type.
type Int32Reader interface {
	Read() (int32, error)
}

// Int32LimitReader defines an interface for reading a slice of int32 type.
type Int32LimitReader interface {
	ReadLimit(int) ([]int32, error)
}

// Int32ReadCloser defines an interface for reading a single int32 type.
type Int32ReadCloser interface {
	Closer
	Int32Reader
}

// Int32Writer defines an interface for writing a single int32 type.
type Int32Writer interface {
	Write(int32) (int, error)
}

// Int32LimitWriter defines an interface for writing a slice of int32 type.
type Int32LimitWriter interface {
	WriteLimit([]int32) (int, error)
}

// Int32WriteCloser defines an interface for writing a single int32 type.
type Int32WriteCloser interface {
	Closer
	Int32Writer
}
