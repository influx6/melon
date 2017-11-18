package melon

import (
	"github.com/influx6/faux/context"
)

// ByteUniqueHash defines a unique hash for Byte which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const ByteUniqueHash = "9f2e300d92751ef08344907e4adc0481d7f9ae93"

// ByteUnitReaderFunc defines a function which expects the giving ByteReader type has input.
type ByteUnitReaderFunc func(ByteReader) ByteUnitReader

// ByteUnitReaderFuncWithContext defines a function which expects the giving ByteReader type has input.
// This expects to receive a context.Context type.
type ByteUnitReaderFuncWithContext func(context.Context, ByteReader) ByteUnitReader

// ByteReaderFunc defines a function which expects the giving ByteReader type has input.
type ByteReaderFunc func(ByteReader) error

// ByteReaderFuncWithContext defines a function which expects the giving ByteReader type has input.
// This expects to receive a context.Context type.
type ByteReaderFuncWithContext func(context.Context, ByteReader) error

// ByteReadCloserFunc defines a function which expects the giving ByteReadCloser type has input.
type ByteReadCloserFunc func(ByteReadCloser) error

// ByteReadCloserFuncWithContext defines a function which expects the giving ByteReadCloser type has input.
// This expects to receive a context.Context type.
type ByteReadCloserFuncWithContext func(context.Context, ByteReadCloser) error

// ByteWriterFunc defines a function which expects the giving ByteWriter type has input.
type ByteWriterFunc func(ByteWriter) error

// ByteWriteCloserFunc defines a function which expects the giving ByteWriteCloser type has input.
type ByteWriteCloserFunc func(ByteWriteCloser) error

// ByteWriterFuncWithContext defines a function which expects the giving ByteWriter type has input.
// This expects to receive a context.Context type.
type ByteWriterFuncWithContext func(context.Context, ByteWriter) error

// ByteWriteCloserFuncWithContext defines a function which expects the giving ByteWriteCloser type has input.
// This expects to receive a context.Context type.
type ByteWriteCloserFuncWithContext func(context.Context, ByteWriteCloser) error

// ByteReader defines an interface for reading a slice of byte types.
type ByteReader interface {
	Read([]byte) (int, error)
}

// ByteReadCloser defines an interface for reading a slice of byte types.
type ByteReadCloser interface {
	Closer
	ByteReader
}

// ByteUnitReader defines an interface for reading a single item of byte type.
type ByteUnitReader interface {
	ReadUnit() (byte, error)
}

// ByteWriter defines an interface for writing a slice of byte types.
type ByteWriter interface {
	Write([]byte) (int, error)
}

// ByteUnitWriter defines an interface for writing a single byte type.
type ByteUnitWriter interface {
	WriteUnit(byte) (int, error)
}

// ByteWriteCloser defines an interface for writing a slice of byte types.
type ByteWriteCloser interface {
	Closer
	ByteWriter
}
