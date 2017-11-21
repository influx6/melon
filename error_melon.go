package melon

// ErrorUniqueHash defines a unique hash for Error which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const ErrorUniqueHash = "d3e8aa638ed54763fe05bd5404e80abdfc12b4f5"

// ErrorReader defines an interface for reading a single error type.
type ErrorReader interface {
	Read() (error, error)
}

// ErrorLimitReader defines an interface for reading a slice of error type.
type ErrorLimitReader interface {
	ReadLimit(int) ([]error, error)
}

// ErrorReadCloser defines an interface for reading a single error type.
type ErrorReadCloser interface {
	Closer
	ErrorReader
}

// ErrorWriter defines an interface for writing a single error type.
type ErrorWriter interface {
	Write(error) (int, error)
}

// ErrorLimitWriter defines an interface for writing a slice of error type.
type ErrorLimitWriter interface {
	WriteLimit([]error) (int, error)
}

// ErrorWriteCloser defines an interface for writing a single error type.
type ErrorWriteCloser interface {
	Closer
	ErrorWriter
}
