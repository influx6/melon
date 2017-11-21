package melon

// Float64UniqueHash defines a unique hash for Float64 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Float64UniqueHash = "c11c63ba6f4639c386f2b6ed2ec7dec76ff35a5d"

// Float64Reader defines an interface for reading a single float64 type.
type Float64Reader interface {
	Read() (float64, error)
}

// Float64LimitReader defines an interface for reading a slice of float64 type.
type Float64LimitReader interface {
	ReadLimit(int) ([]float64, error)
}

// Float64ReadCloser defines an interface for reading a single float64 type.
type Float64ReadCloser interface {
	Closer
	Float64Reader
}

// Float64Writer defines an interface for writing a single float64 type.
type Float64Writer interface {
	Write(float64) (int, error)
}

// Float64LimitWriter defines an interface for writing a slice of float64 type.
type Float64LimitWriter interface {
	WriteLimit([]float64) (int, error)
}

// Float64WriteCloser defines an interface for writing a single float64 type.
type Float64WriteCloser interface {
	Closer
	Float64Writer
}
