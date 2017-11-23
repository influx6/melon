import (
  "github.com/influx6/faux/context"
  {{range attrs "imports"}}
  {{quote .}}
  {{end}}
)

// {{sel "Name" | capitalize}}UniqueHash defines a unique hash for {{sel "Name"}} which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const {{sel "Name" | capitalize}}UniqueHash = {{ (joinVariadic ":" (sel "Name") (sel "Type")) | sha1 | quote }}

// {{sel "Name"}}Handler defines a function type receiving both reader and writer types.
type {{sel "Name"}}Handler func({{sel "Name"}}Reader, {{sel "Name"}}Writer) error

// {{sel "Name"}}HandlerWithCtx defines a function type receiving both reader and writer types.
type {{sel "Name"}}HandlerWithCtx func(context.Context,{{sel "Name"}}Reader, {{sel "Name"}}Writer) error

// {{sel "Name"}}Stream defines a function type receiving both reader and writer types.
type {{sel "Name"}}Stream func({{sel "Name"}}StreamReader, {{sel "Name"}}StreamWriter) error

// {{sel "Name"}}StreamWithCtx defines a function type receiving both reader and writer types.
type {{sel "Name"}}StreamWithCtx func(context.Context,{{sel "Name"}}StreamReader, {{sel "Name"}}StreamWriter) error

// {{sel "Name"}}Reader defines an interface for reading a single {{sel "Type"}} type.
type {{sel "Name"}}Reader interface {
  Read{{sel "Name"}}() ({{sel "Type"}}, error)
}

// {{sel "Name"}}ReadCloser defines an interface for reading a single {{sel "Type"}} type.
type {{sel "Name"}}ReadCloser interface {
  Closer
  {{sel "Name"}}Reader
}

// {{sel "Name"}}StreamReader defines an interface for reading a slice of {{sel "Type"}} type.
type {{sel "Name"}}StreamReader interface {
  Read(int) ([]{{sel "Type"}}, error)
}

// {{sel "Name"}}StreamReadCloser defines an interface for reading a single {{sel "Type"}} type.
type {{sel "Name"}}StreamReadCloser interface {
  Closer
  {{sel "Name"}}StreamReader
}

// {{sel "Name"}}Writer defines an interface for writing a single {{sel "Type"}} type.
type {{sel "Name"}}Writer interface {
  Write{{sel "Name"}}({{sel "Type"}}) (int, error)
}

// {{sel "Name"}}WriteCloser defines an interface for writing a single {{sel "Type"}} type.
type {{sel "Name"}}WriteCloser interface {
  Closer
  {{sel "Name"}}Writer
}

// {{sel "Name"}}StreamWriter defines an interface for writing a slice of {{sel "Type"}} type.
type {{sel "Name"}}StreamWriter interface {
  Write([]{{sel "Type"}}) (int, error)
}

// {{sel "Name"}}StreamWriteCloser defines an interface for writing a single {{sel "Type"}} type.
type {{sel "Name"}}StreamWriteCloser interface {
  Closer
  {{sel "Name"}}StreamWriter
}
