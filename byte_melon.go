package melon

// ByteUniqueHash defines a unique hash for Byte which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const ByteUniqueHash = "9f2e300d92751ef08344907e4adc0481d7f9ae93"

// ByteReader defines an interface for reading a single byte type.
type ByteReader interface {
	Read() (byte, error)
}

// ByteLimitReader defines an interface for reading a slice of byte type.
type ByteLimitReader interface {
	ReadLimit(int) ([]byte, error)
}

// ByteReadCloser defines an interface for reading a single byte type.
type ByteReadCloser interface {
	Closer
	ByteReader
}

// ByteWriter defines an interface for writing a single byte type.
type ByteWriter interface {
	Write(byte) (int, error)
}

// ByteLimitWriter defines an interface for writing a slice of byte type.
type ByteLimitWriter interface {
	WriteLimit([]byte) (int, error)
}

// ByteWriteCloser defines an interface for writing a single byte type.
type ByteWriteCloser interface {
	Closer
	ByteWriter
}
