package reading

import "go/types"

type Store struct {
	Models  map[string]*types.Named
	Read    bool
	Written bool
	Calls   map[string]bool
}
