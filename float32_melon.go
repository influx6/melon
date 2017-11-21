package melon

// Float32UniqueHash defines a unique hash for Float32 which can
// be used to reference a given instance within a context.ValueBag or a google context.Context
// value store.
const Float32UniqueHash = "c7cd93eba0df834db18da1b14dc29a7551a1dd80"

// Float32Reader defines an interface for reading a single float32 type.
type Float32Reader interface {
	Read() (float32, error)
}

// Float32LimitReader defines an interface for reading a slice of float32 type.
type Float32LimitReader interface {
	ReadLimit(int) ([]float32, error)
}

// Float32ReadCloser defines an interface for reading a single float32 type.
type Float32ReadCloser interface {
	Closer
	Float32Reader
}

// Float32Writer defines an interface for writing a single float32 type.
type Float32Writer interface {
	Write(float32) (int, error)
}

// Float32LimitWriter defines an interface for writing a slice of float32 type.
type Float32LimitWriter interface {
	WriteLimit([]float32) (int, error)
}

// Float32WriteCloser defines an interface for writing a single float32 type.
type Float32WriteCloser interface {
	Closer
	Float32Writer
}
