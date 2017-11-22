package melon

// ErrorUniqueHash defines a unique hash for Error which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const ErrorUniqueHash = "d3e8aa638ed54763fe05bd5404e80abdfc12b4f5"

// ErrorReader defines an interface for reading a single error type.
type ErrorReader interface {
	ReadError() (error, error)
}

// ErrorReadCloser defines an interface for reading a single error type.
type ErrorReadCloser interface {
	Closer
	ErrorReader
}

// ErrorStreamReader defines an interface for reading a slice of error type.
type ErrorStreamReader interface {
	Read(int) ([]error, error)
}

// ErrorStreamReadCloser defines an interface for reading a single error type.
type ErrorStreamReadCloser interface {
	Closer
	ErrorStreamReader
}

// ErrorWriter defines an interface for writing a single error type.
type ErrorWriter interface {
	WriteError(error) (int, error)
}

// ErrorWriteCloser defines an interface for writing a single error type.
type ErrorWriteCloser interface {
	Closer
	ErrorWriter
}

// ErrorStreamWriter defines an interface for writing a slice of error type.
type ErrorStreamWriter interface {
	Write([]error) (int, error)
}

// ErrorStreamWriteCloser defines an interface for writing a single error type.
type ErrorStreamWriteCloser interface {
	Closer
	ErrorStreamWriter
}
