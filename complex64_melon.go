package melon

// Complex64UniqueHash defines a unique hash for Complex64 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Complex64UniqueHash = "161e6337d01e5b8bd4e79ac4060fde11263b2d17"

// Complex64Reader defines an interface for reading a single complex64 type.
type Complex64Reader interface {
	Read() (complex64, error)
}

// Complex64LimitReader defines an interface for reading a slice of complex64 type.
type Complex64LimitReader interface {
	ReadLimit(int) ([]complex64, error)
}

// Complex64ReadCloser defines an interface for reading a single complex64 type.
type Complex64ReadCloser interface {
	Closer
	Complex64Reader
}

// Complex64Writer defines an interface for writing a single complex64 type.
type Complex64Writer interface {
	Write(complex64) (int, error)
}

// Complex64LimitWriter defines an interface for writing a slice of complex64 type.
type Complex64LimitWriter interface {
	WriteLimit([]complex64) (int, error)
}

// Complex64WriteCloser defines an interface for writing a single complex64 type.
type Complex64WriteCloser interface {
	Closer
	Complex64Writer
}
