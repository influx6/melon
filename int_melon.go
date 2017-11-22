package melon

// IntUniqueHash defines a unique hash for Int which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const IntUniqueHash = "e3e6a2908db2b977524568702fabd849a5d5514f"

// IntReader defines an interface for reading a single int type.
type IntReader interface {
	ReadInt() (int, error)
}

// IntReadCloser defines an interface for reading a single int type.
type IntReadCloser interface {
	Closer
	IntReader
}

// IntStreamReader defines an interface for reading a slice of int type.
type IntStreamReader interface {
	Read(int) ([]int, error)
}

// IntStreamReadCloser defines an interface for reading a single int type.
type IntStreamReadCloser interface {
	Closer
	IntStreamReader
}

// IntWriter defines an interface for writing a single int type.
type IntWriter interface {
	WriteInt(int) (int, error)
}

// IntWriteCloser defines an interface for writing a single int type.
type IntWriteCloser interface {
	Closer
	IntWriter
}

// IntStreamWriter defines an interface for writing a slice of int type.
type IntStreamWriter interface {
	Write([]int) (int, error)
}

// IntStreamWriteCloser defines an interface for writing a single int type.
type IntStreamWriteCloser interface {
	Closer
	IntStreamWriter
}
