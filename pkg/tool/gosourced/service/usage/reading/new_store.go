package reading

import "go/types"

func NewStore() *Store {
	return &Store{Models: map[string]*types.Named{}, Calls: map[string]bool{}}
}
