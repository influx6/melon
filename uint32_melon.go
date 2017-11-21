package melon

// UInt32UniqueHash defines a unique hash for UInt32 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const UInt32UniqueHash = "a334cf504c1433f88e8aaab77d2225c3f7843634"

// UInt32Reader defines an interface for reading a single uint32 type.
type UInt32Reader interface {
	Read() (uint32, error)
}

// UInt32LimitReader defines an interface for reading a slice of uint32 type.
type UInt32LimitReader interface {
	ReadLimit(int) ([]uint32, error)
}

// UInt32ReadCloser defines an interface for reading a single uint32 type.
type UInt32ReadCloser interface {
	Closer
	UInt32Reader
}

// UInt32Writer defines an interface for writing a single uint32 type.
type UInt32Writer interface {
	Write(uint32) (int, error)
}

// UInt32LimitWriter defines an interface for writing a slice of uint32 type.
type UInt32LimitWriter interface {
	WriteLimit([]uint32) (int, error)
}

// UInt32WriteCloser defines an interface for writing a single uint32 type.
type UInt32WriteCloser interface {
	Closer
	UInt32Writer
}
