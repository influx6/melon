package melon

// UInt8UniqueHash defines a unique hash for UInt8 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const UInt8UniqueHash = "e5de7d77d24ebbe134934f7c5f8dc3fafb344aba"

// UInt8Reader defines an interface for reading a single uint8 type.
type UInt8Reader interface {
	ReadUInt8() (uint8, error)
}

// UInt8ReadCloser defines an interface for reading a single uint8 type.
type UInt8ReadCloser interface {
	Closer
	UInt8Reader
}

// UInt8StreamReader defines an interface for reading a slice of uint8 type.
type UInt8StreamReader interface {
	Read(int) ([]uint8, error)
}

// UInt8StreamReadCloser defines an interface for reading a single uint8 type.
type UInt8StreamReadCloser interface {
	Closer
	UInt8StreamReader
}

// UInt8Writer defines an interface for writing a single uint8 type.
type UInt8Writer interface {
	WriteUInt8(uint8) (int, error)
}

// UInt8WriteCloser defines an interface for writing a single uint8 type.
type UInt8WriteCloser interface {
	Closer
	UInt8Writer
}

// UInt8StreamWriter defines an interface for writing a slice of uint8 type.
type UInt8StreamWriter interface {
	Write([]uint8) (int, error)
}

// UInt8StreamWriteCloser defines an interface for writing a single uint8 type.
type UInt8StreamWriteCloser interface {
	Closer
	UInt8StreamWriter
}
