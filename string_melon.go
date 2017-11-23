package melon

import (
	"github.com/influx6/faux/context"
)

// StringUniqueHash defines a unique hash for String which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const StringUniqueHash = "62a027764a58e833c7beb34a33ccb1584d611d17"

// StringHandler defines a function type receiving both reader and writer types.
type StringHandler func(StringReader, StringWriter) error

// StringHandlerWithCtx defines a function type receiving both reader and writer types.
type StringHandlerWithCtx func(context.Context, StringReader, StringWriter) error

// StringStream defines a function type receiving both reader and writer types.
type StringStream func(StringStreamReader, StringStreamWriter) error

// StringStreamWithCtx defines a function type receiving both reader and writer types.
type StringStreamWithCtx func(context.Context, StringStreamReader, StringStreamWriter) error

// StringReader defines an interface for reading a single string type.
type StringReader interface {
	ReadString() (string, error)
}

// StringReadCloser defines an interface for reading a single string type.
type StringReadCloser interface {
	Closer
	StringReader
}

// StringStreamReader defines an interface for reading a slice of string type.
type StringStreamReader interface {
	Read(int) ([]string, error)
}

// StringStreamReadCloser defines an interface for reading a single string type.
type StringStreamReadCloser interface {
	Closer
	StringStreamReader
}

// StringWriter defines an interface for writing a single string type.
type StringWriter interface {
	WriteString(string) (int, error)
}

// StringWriteCloser defines an interface for writing a single string type.
type StringWriteCloser interface {
	Closer
	StringWriter
}

// StringStreamWriter defines an interface for writing a slice of string type.
type StringStreamWriter interface {
	Write([]string) (int, error)
}

// StringStreamWriteCloser defines an interface for writing a single string type.
type StringStreamWriteCloser interface {
	Closer
	StringStreamWriter
}
