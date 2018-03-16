package storage

// Getter gets module metadata and its source from underlying storage
type Getter interface {
	// must return NotFoundErr if the coordinates are not found
	Get(module, vsn string) (*Version, error)
}
