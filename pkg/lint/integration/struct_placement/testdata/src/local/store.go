package local

type Store struct{}

func (s *Store) Save() {}

func (s *Store) local() {
	type row struct {
		name string
	}
	_ = row{}
}
