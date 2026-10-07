package tagged

import "github.com/funtimecoding/soil/pkg/console/types/raw"

func New(
	identifier int,
	name string,
	rawName string,
) *Tagged {
	return &Tagged{
		Identifier: identifier,
		Name:       name,
		Raw:        raw.New(rawName),
	}
}
