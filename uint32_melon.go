package melon

import (
	"github.com/influx6/faux/context"
)

// UInt32UnitReaderFunc defines a function which expects the giving UInt32Reader type has input.
type UInt32UnitReaderFunc func(UInt32Reader) UInt32UnitReader

// UInt32UnitReaderFuncWithContext defines a function which expects the giving UInt32Reader type has input.
// This expects to receive a context.Context type.
type UInt32UnitReaderFuncWithContext func(context.Context, UInt32Reader) UInt32UnitReader

// UInt32ReaderFunc defines a function which expects the giving UInt32Reader type has input.
type UInt32ReaderFunc func(UInt32Reader) error

// UInt32ReaderFuncWithContext defines a function which expects the giving UInt32Reader type has input.
// This expects to receive a context.Context type.
type UInt32ReaderFuncWithContext func(context.Context, UInt32Reader) error

// UInt32ReadCloserFunc defines a function which expects the giving UInt32ReadCloser type has input.
type UInt32ReadCloserFunc func(UInt32ReadCloser) error

// UInt32ReadCloserFuncWithContext defines a function which expects the giving UInt32ReadCloser type has input.
// This expects to receive a context.Context type.
type UInt32ReadCloserFuncWithContext func(context.Context, UInt32ReadCloser) error

// UInt32WriterFunc defines a function which expects the giving UInt32Writer type has input.
type UInt32WriterFunc func(UInt32Writer) error

// UInt32WriteCloserFunc defines a function which expects the giving UInt32WriteCloser type has input.
type UInt32WriteCloserFunc func(UInt32WriteCloser) error

// UInt32WriterFuncWithContext defines a function which expects the giving UInt32Writer type has input.
// This expects to receive a context.Context type.
type UInt32WriterFuncWithContext func(context.Context, UInt32Writer) error

// UInt32WriteCloserFuncWithContext defines a function which expects the giving UInt32WriteCloser type has input.
// This expects to receive a context.Context type.
type UInt32WriteCloserFuncWithContext func(context.Context, UInt32WriteCloser) error

// UInt32Reader defines an interface for reading a slice of uint32 types.
type UInt32Reader interface {
	Read([]uint32) (int, error)
}

// UInt32ReadCloser defines an interface for reading a slice of uint32 types.
type UInt32ReadCloser interface {
	Closer
	UInt32Reader
}

// UInt32UnitReader defines an interface for reading a single item of uint32 type.
type UInt32UnitReader interface {
	ReadUnit() (uint32, error)
}

// UInt32Writer defines an interface for writing a slice of uint32 types.
type UInt32Writer interface {
	Write([]uint32) (int, error)
}

// UInt32UnitWriter defines an interface for writing a single uint32 type.
type UInt32UnitWriter interface {
	WriteUnit(uint32) (int, error)
}

// UInt32WriteCloser defines an interface for writing a slice of uint32 types.
type UInt32WriteCloser interface {
	Closer
	UInt32Writer
}
