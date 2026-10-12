package store_tester

func New() *Tester {
	return &Tester{
		stored:   map[string]map[string]string{},
		existing: map[string]string{},
	}
}
