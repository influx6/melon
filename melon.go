package melon

// Closer defines a inteface which exposes a single Close method.
type Closer interface {
	Close() error
}

// Seeker defines a interface which exposes a single Seek method.
type Seeker interface {
	Seek(int) error
}
