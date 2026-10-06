package references

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/location"

func New(
	symbol string,
	locations []*location.Location,
) *References {
	return &References{
		Symbol:    symbol,
		Total:     len(locations),
		Locations: locations,
	}
}
