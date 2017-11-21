package melon

import "net"

// ConnUniqueHash defines a unique hash for Conn which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const ConnUniqueHash = "e227c185e376f84717e05beb52bde64ab47d6838"

// ConnReader defines an interface for reading a single net.Conn type.
type ConnReader interface {
	Read() (net.Conn, error)
}

// ConnLimitReader defines an interface for reading a slice of net.Conn type.
type ConnLimitReader interface {
	ReadLimit(int) ([]net.Conn, error)
}

// ConnReadCloser defines an interface for reading a single net.Conn type.
type ConnReadCloser interface {
	Closer
	ConnReader
}

// ConnWriter defines an interface for writing a single net.Conn type.
type ConnWriter interface {
	Write(net.Conn) (int, error)
}

// ConnLimitWriter defines an interface for writing a slice of net.Conn type.
type ConnLimitWriter interface {
	WriteLimit([]net.Conn) (int, error)
}

// ConnWriteCloser defines an interface for writing a single net.Conn type.
type ConnWriteCloser interface {
	Closer
	ConnWriter
}
