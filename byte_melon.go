package melon

// ByteUniqueHash defines a unique hash for Byte which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const ByteUniqueHash = "9f2e300d92751ef08344907e4adc0481d7f9ae93"

// ByteReader defines an interface for reading a single byte type.
type ByteReader interface {
	ReadByte() (byte, error)
}

// ByteReadCloser defines an interface for reading a single byte type.
type ByteReadCloser interface {
	Closer
	ByteReader
}

// ByteStreamReader defines an interface for reading a slice of byte type.
type ByteStreamReader interface {
	Read(int) ([]byte, error)
}

// ByteStreamReadCloser defines an interface for reading a single byte type.
type ByteStreamReadCloser interface {
	Closer
	ByteStreamReader
}

// ByteWriter defines an interface for writing a single byte type.
type ByteWriter interface {
	WriteByte(byte) (int, error)
}

// ByteWriteCloser defines an interface for writing a single byte type.
type ByteWriteCloser interface {
	Closer
	ByteWriter
}

// ByteStreamWriter defines an interface for writing a slice of byte type.
type ByteStreamWriter interface {
	Write([]byte) (int, error)
}

// ByteStreamWriteCloser defines an interface for writing a single byte type.
type ByteStreamWriteCloser interface {
	Closer
	ByteStreamWriter
}
