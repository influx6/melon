package melon

// UInt64UniqueHash defines a unique hash for UInt64 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const UInt64UniqueHash = "9dfe93d26e38e94f8348337651317c3130bb021d"

// UInt64Reader defines an interface for reading a single uint64 type.
type UInt64Reader interface {
	Read() (uint64, error)
}

// UInt64LimitReader defines an interface for reading a slice of uint64 type.
type UInt64LimitReader interface {
	ReadLimit(int) ([]uint64, error)
}

// UInt64ReadCloser defines an interface for reading a single uint64 type.
type UInt64ReadCloser interface {
	Closer
	UInt64Reader
}

// UInt64Writer defines an interface for writing a single uint64 type.
type UInt64Writer interface {
	Write(uint64) (int, error)
}

// UInt64LimitWriter defines an interface for writing a slice of uint64 type.
type UInt64LimitWriter interface {
	WriteLimit([]uint64) (int, error)
}

// UInt64WriteCloser defines an interface for writing a single uint64 type.
type UInt64WriteCloser interface {
	Closer
	UInt64Writer
}
