import (
  "github.com/influx6/faux/context"
)

// {{sel "Name" | capitalize}}UniqueHash defines a unique hash for {{sel "Name"}} which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const {{sel "Name" | capitalize}}UniqueHash = {{ (joinVariadic ":" (sel "Name") (sel "Type")) | sha1 | quote }}

// {{sel "Name"}}UnitReaderFunc defines a function which expects the giving {{sel "Name"}}Reader type has input.
type {{sel "Name"}}UnitReaderFunc func({{sel "Name"}}Reader) {{sel "Name"}}UnitReader

// {{sel "Name"}}UnitReaderFuncWithContext defines a function which expects the giving {{sel "Name"}}Reader type has input.
// This expects to receive a context.Context type.
type {{sel "Name"}}UnitReaderFuncWithContext func(context.Context, {{sel "Name"}}Reader) {{sel "Name"}}UnitReader

// {{sel "Name"}}ReaderFunc defines a function which expects the giving {{sel "Name"}}Reader type has input.
type {{sel "Name"}}ReaderFunc func({{sel "Name"}}Reader) error

// {{sel "Name"}}ReaderFuncWithContext defines a function which expects the giving {{sel "Name"}}Reader type has input.
// This expects to receive a context.Context type.
type {{sel "Name"}}ReaderFuncWithContext func(context.Context,{{sel "Name"}}Reader) error

// {{sel "Name"}}ReadCloserFunc defines a function which expects the giving {{sel "Name"}}ReadCloser type has input.
type {{sel "Name"}}ReadCloserFunc func({{sel "Name"}}ReadCloser) error

// {{sel "Name"}}ReadCloserFuncWithContext defines a function which expects the giving {{sel "Name"}}ReadCloser type has input.
// This expects to receive a context.Context type.
type {{sel "Name"}}ReadCloserFuncWithContext func(context.Context,{{sel "Name"}}ReadCloser) error

// {{sel "Name"}}WriterFunc defines a function which expects the giving {{sel "Name"}}Writer type has input.
type {{sel "Name"}}WriterFunc func({{sel "Name"}}Writer) error

// {{sel "Name"}}WriteCloserFunc defines a function which expects the giving {{sel "Name"}}WriteCloser type has input.
type {{sel "Name"}}WriteCloserFunc func({{sel "Name"}}WriteCloser) error

// {{sel "Name"}}WriterFuncWithContext defines a function which expects the giving {{sel "Name"}}Writer type has input.
// This expects to receive a context.Context type.
type {{sel "Name"}}WriterFuncWithContext func(context.Context,{{sel "Name"}}Writer) error

// {{sel "Name"}}WriteCloserFuncWithContext defines a function which expects the giving {{sel "Name"}}WriteCloser type has input.
// This expects to receive a context.Context type.
type {{sel "Name"}}WriteCloserFuncWithContext func(context.Context,{{sel "Name"}}WriteCloser) error

// {{sel "Name"}}Reader defines an interface for reading a slice of {{sel "Type"}} types.
type {{sel "Name"}}Reader interface {
  Read([]{{sel "Type"}}) (int, error)
}

// {{sel "Name"}}ReadCloser defines an interface for reading a slice of {{sel "Type"}} types.
type {{sel "Name"}}ReadCloser interface {
  Closer
  {{sel "Name"}}Reader
}

// {{sel "Name"}}UnitReader defines an interface for reading a single item of {{sel "Type"}} type.
type {{sel "Name"}}UnitReader interface {
  ReadUnit() ({{sel "Type"}}, error)
}

// {{sel "Name"}}Writer defines an interface for writing a slice of {{sel "Type"}} types.
type {{sel "Name"}}Writer interface {
  Write([]{{sel "Type"}}) (int, error)
}

// {{sel "Name"}}UnitWriter defines an interface for writing a single {{sel "Type"}} type.
type {{sel "Name"}}UnitWriter interface {
  WriteUnit({{sel "Type"}}) (int, error)
}

// {{sel "Name"}}WriteCloser defines an interface for writing a slice of {{sel "Type"}} types.
type {{sel "Name"}}WriteCloser interface {
  Closer
  {{sel "Name"}}Writer
}
