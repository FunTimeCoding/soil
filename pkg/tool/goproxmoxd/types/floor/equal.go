package floor

import "slices"

func (f Floor) Equal(other Floor) bool {
	return slices.Equal(f.Nodes, other.Nodes) &&
		slices.Equal(f.Guests, other.Guests) &&
		slices.Equal(f.Storages, other.Storages)
}
