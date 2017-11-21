package melon

// StringUniqueHash defines a unique hash for String which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const StringUniqueHash = "62a027764a58e833c7beb34a33ccb1584d611d17"

// StringReader defines an interface for reading a single string type.
type StringReader interface {
	Read() (string, error)
}

// StringLimitReader defines an interface for reading a slice of string type.
type StringLimitReader interface {
	ReadLimit(int) ([]string, error)
}

// StringReadCloser defines an interface for reading a single string type.
type StringReadCloser interface {
	Closer
	StringReader
}

// StringWriter defines an interface for writing a single string type.
type StringWriter interface {
	Write(string) (int, error)
}

// StringLimitWriter defines an interface for writing a slice of string type.
type StringLimitWriter interface {
	WriteLimit([]string) (int, error)
}

// StringWriteCloser defines an interface for writing a single string type.
type StringWriteCloser interface {
	Closer
	StringWriter
}
