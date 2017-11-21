package melon

// UInt8UniqueHash defines a unique hash for UInt8 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const UInt8UniqueHash = "e5de7d77d24ebbe134934f7c5f8dc3fafb344aba"

// UInt8Reader defines an interface for reading a single uint8 type.
type UInt8Reader interface {
	Read() (uint8, error)
}

// UInt8LimitReader defines an interface for reading a slice of uint8 type.
type UInt8LimitReader interface {
	ReadLimit(int) ([]uint8, error)
}

// UInt8ReadCloser defines an interface for reading a single uint8 type.
type UInt8ReadCloser interface {
	Closer
	UInt8Reader
}

// UInt8Writer defines an interface for writing a single uint8 type.
type UInt8Writer interface {
	Write(uint8) (int, error)
}

// UInt8LimitWriter defines an interface for writing a slice of uint8 type.
type UInt8LimitWriter interface {
	WriteLimit([]uint8) (int, error)
}

// UInt8WriteCloser defines an interface for writing a single uint8 type.
type UInt8WriteCloser interface {
	Closer
	UInt8Writer
}
