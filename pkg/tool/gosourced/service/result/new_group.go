package result

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/location"

func NewGroup(
	shape string,
	exemplar string,
	locations []*location.Location,
) *Group {
	return &Group{Shape: shape, Exemplar: exemplar, Locations: locations}
}
