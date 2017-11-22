package melon

import "net"

// ConnUniqueHash defines a unique hash for Conn which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const ConnUniqueHash = "e227c185e376f84717e05beb52bde64ab47d6838"

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
