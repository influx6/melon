package melon

import (
	"github.com/influx6/faux/context"
)

// UInt8UniqueHash defines a unique hash for UInt8 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const UInt8UniqueHash = "e5de7d77d24ebbe134934f7c5f8dc3fafb344aba"

// UInt8UnitReaderFunc defines a function which expects the giving UInt8Reader type has input.
type UInt8UnitReaderFunc func(UInt8Reader) UInt8UnitReader

// UInt8UnitReaderFuncWithContext defines a function which expects the giving UInt8Reader type has input.
// This expects to receive a context.Context type.
type UInt8UnitReaderFuncWithContext func(context.Context, UInt8Reader) UInt8UnitReader

// UInt8ReaderFunc defines a function which expects the giving UInt8Reader type has input.
type UInt8ReaderFunc func(UInt8Reader) error

// UInt8ReaderFuncWithContext defines a function which expects the giving UInt8Reader type has input.
// This expects to receive a context.Context type.
type UInt8ReaderFuncWithContext func(context.Context, UInt8Reader) error

// UInt8ReadCloserFunc defines a function which expects the giving UInt8ReadCloser type has input.
type UInt8ReadCloserFunc func(UInt8ReadCloser) error

// UInt8ReadCloserFuncWithContext defines a function which expects the giving UInt8ReadCloser type has input.
// This expects to receive a context.Context type.
type UInt8ReadCloserFuncWithContext func(context.Context, UInt8ReadCloser) error

// UInt8WriterFunc defines a function which expects the giving UInt8Writer type has input.
type UInt8WriterFunc func(UInt8Writer) error

// UInt8WriteCloserFunc defines a function which expects the giving UInt8WriteCloser type has input.
type UInt8WriteCloserFunc func(UInt8WriteCloser) error

// UInt8WriterFuncWithContext defines a function which expects the giving UInt8Writer type has input.
// This expects to receive a context.Context type.
type UInt8WriterFuncWithContext func(context.Context, UInt8Writer) error

// UInt8WriteCloserFuncWithContext defines a function which expects the giving UInt8WriteCloser type has input.
// This expects to receive a context.Context type.
type UInt8WriteCloserFuncWithContext func(context.Context, UInt8WriteCloser) error

// UInt8Reader defines an interface for reading a slice of uint8 types.
type UInt8Reader interface {
	Read([]uint8) (int, error)
}

// UInt8ReadCloser defines an interface for reading a slice of uint8 types.
type UInt8ReadCloser interface {
	Closer
	UInt8Reader
}

// UInt8UnitReader defines an interface for reading a single item of uint8 type.
type UInt8UnitReader interface {
	ReadUnit() (uint8, error)
}

// UInt8Writer defines an interface for writing a slice of uint8 types.
type UInt8Writer interface {
	Write([]uint8) (int, error)
}

// UInt8UnitWriter defines an interface for writing a single uint8 type.
type UInt8UnitWriter interface {
	WriteUnit(uint8) (int, error)
}

// UInt8WriteCloser defines an interface for writing a slice of uint8 types.
type UInt8WriteCloser interface {
	Closer
	UInt8Writer
}
