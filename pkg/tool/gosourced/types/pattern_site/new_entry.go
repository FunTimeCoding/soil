package pattern_site

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/result/location"

func NewEntry(
	shape string,
	exemplar string,
	location *location.Location,
) *Entry {
	return &Entry{Shape: shape, Exemplar: exemplar, Location: location}
}
