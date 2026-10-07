package xref

import "github.com/funtimecoding/soil/pkg/source/index/record"

func New(
	units map[string]*record.References,
	directories map[string]string,
) *Index {
	return &Index{units: units, directories: directories}
}
