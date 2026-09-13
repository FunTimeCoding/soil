package service

import "github.com/funtimecoding/soil/pkg/directory"

func first(
	record *directory.Record,
	name string,
) string {
	values := record.Attributes[name]

	if len(values) == 0 {
		return ""
	}

	return values[0]
}
