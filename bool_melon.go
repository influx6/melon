package melon

// BoolUniqueHash defines a unique hash for Bool which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const BoolUniqueHash = "189fa523325f8a701fd00ad1b0cd386b4b629299"

// BoolReader defines an interface for reading a single bool type.
type BoolReader interface {
	Read() (bool, error)
}

// BoolLimitReader defines an interface for reading a slice of bool type.
type BoolLimitReader interface {
	ReadLimit(int) ([]bool, error)
}

// BoolReadCloser defines an interface for reading a single bool type.
type BoolReadCloser interface {
	Closer
	BoolReader
}

// BoolWriter defines an interface for writing a single bool type.
type BoolWriter interface {
	Write(bool) (int, error)
}

// BoolLimitWriter defines an interface for writing a slice of bool type.
type BoolLimitWriter interface {
	WriteLimit([]bool) (int, error)
}

// BoolWriteCloser defines an interface for writing a single bool type.
type BoolWriteCloser interface {
	Closer
	BoolWriter
}
