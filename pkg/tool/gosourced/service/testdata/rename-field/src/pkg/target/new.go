package target

func New(api string) *Store {
	return &Store{api: api}
}
