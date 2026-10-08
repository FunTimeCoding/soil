package local

type Store struct{}

func (s *Store) Save() {
	type row struct {
		name string
	}
	_ = row{}
}
