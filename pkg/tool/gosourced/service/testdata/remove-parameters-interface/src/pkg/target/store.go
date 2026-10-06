package target

type Saver interface {
	Save(name string, version string) error
}

type Store struct{}

func (s *Store) Save(
	name string,
	version string,
) error {
	return nil
}
