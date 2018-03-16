package storage

// Lister is the interface that lists versions of a specific baseURL & module
type Lister interface {
	// List gets all the versions for the given baseURL & module.
	// It returns NotFoundErr if baseURL/module isn't found
	List(module string) ([]string, error)
}
