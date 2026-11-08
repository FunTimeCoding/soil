package census

import "go/types"

type storeReading struct {
	models  map[string]*types.Named
	read    bool
	written bool
	calls   map[string]bool
}
