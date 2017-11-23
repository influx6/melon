package melon

import (
	"net"

	"github.com/influx6/faux/context"
)

// ConnUniqueHash defines a unique hash for Conn which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const ConnUniqueHash = "e227c185e376f84717e05beb52bde64ab47d6838"

// ConnHandler defines a function type receiving both reader and writer types.
type ConnHandler func(ConnReader, ConnWriter) error

// ConnHandlerWithCtx defines a function type receiving both reader and writer types.
type ConnHandlerWithCtx func(context.Context, ConnReader, ConnWriter) error

// ConnStream defines a function type receiving both reader and writer types.
type ConnStream func(ConnStreamReader, ConnStreamWriter) error

// ConnStreamWithCtx defines a function type receiving both reader and writer types.
type ConnStreamWithCtx func(context.Context, ConnStreamReader, ConnStreamWriter) error

// ConnReader defines an interface for reading a single net.Conn type.
type ConnReader interface {
	ReadConn() (net.Conn, error)
}

// ConnReadCloser defines an interface for reading a single net.Conn type.
type ConnReadCloser interface {
	Closer
	ConnReader
}

// ConnStreamReader defines an interface for reading a slice of net.Conn type.
type ConnStreamReader interface {
	Read(int) ([]net.Conn, error)
}

// ConnStreamReadCloser defines an interface for reading a single net.Conn type.
type ConnStreamReadCloser interface {
	Closer
	ConnStreamReader
}

// ConnWriter defines an interface for writing a single net.Conn type.
type ConnWriter interface {
	WriteConn(net.Conn) (int, error)
}

// ConnWriteCloser defines an interface for writing a single net.Conn type.
type ConnWriteCloser interface {
	Closer
	ConnWriter
}

// ConnStreamWriter defines an interface for writing a slice of net.Conn type.
type ConnStreamWriter interface {
	Write([]net.Conn) (int, error)
}

// ConnStreamWriteCloser defines an interface for writing a single net.Conn type.
type ConnStreamWriteCloser interface {
	Closer
	ConnStreamWriter
}
