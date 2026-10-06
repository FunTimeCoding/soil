package store

func New(directory string) *Store {
	return &Store{directory: directory}
}
