package census

import "go/types"

func newStoreReading() *storeReading {
	return &storeReading{
		models: map[string]*types.Named{},
		calls:  map[string]bool{},
	}
}
