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

// {{sel "Name"}}Reader defines an interface for reading a single {{sel "Type"}} type.
type {{sel "Name"}}Reader interface {
  Read() ({{sel "Type"}}, error)
}

// {{sel "Name"}}LimitReader defines an interface for reading a slice of {{sel "Type"}} type.
type {{sel "Name"}}LimitReader interface {
  ReadLimit(int) ([]{{sel "Type"}}, error)
}

// {{sel "Name"}}ReadCloser defines an interface for reading a single {{sel "Type"}} type.
type {{sel "Name"}}ReadCloser interface {
  Closer
  {{sel "Name"}}Reader
}

// {{sel "Name"}}Writer defines an interface for writing a single {{sel "Type"}} type.
type {{sel "Name"}}Writer interface {
  Write({{sel "Type"}}) (int, error)
}

// {{sel "Name"}}LimitWriter defines an interface for writing a slice of {{sel "Type"}} type.
type {{sel "Name"}}LimitWriter interface {
  WriteLimit([]{{sel "Type"}}) (int, error)
}

// {{sel "Name"}}WriteCloser defines an interface for writing a single {{sel "Type"}} type.
type {{sel "Name"}}WriteCloser interface {
  Closer
  {{sel "Name"}}Writer
}
