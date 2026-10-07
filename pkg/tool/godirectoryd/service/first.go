package service

import "github.com/funtimecoding/soil/pkg/directory/types/search_record"

func first(
	record *search_record.Record,
	name string,
) string {
	values := record.Attributes[name]

	if len(values) == 0 {
		return ""
	}

	return values[0]
}
