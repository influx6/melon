package melon

// Complex128UniqueHash defines a unique hash for Complex128 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Complex128UniqueHash = "1e22bb114616c1ea2c888d34d5d519a35cc9103c"

// Complex128Reader defines an interface for reading a single complex128 type.
type Complex128Reader interface {
	Read() (complex128, error)
}

// Complex128LimitReader defines an interface for reading a slice of complex128 type.
type Complex128LimitReader interface {
	ReadLimit(int) ([]complex128, error)
}

// Complex128ReadCloser defines an interface for reading a single complex128 type.
type Complex128ReadCloser interface {
	Closer
	Complex128Reader
}

// Complex128Writer defines an interface for writing a single complex128 type.
type Complex128Writer interface {
	Write(complex128) (int, error)
}

// Complex128LimitWriter defines an interface for writing a slice of complex128 type.
type Complex128LimitWriter interface {
	WriteLimit([]complex128) (int, error)
}

// Complex128WriteCloser defines an interface for writing a single complex128 type.
type Complex128WriteCloser interface {
	Closer
	Complex128Writer
}
