package session_page

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func created(value string) string {
	t, e := time.Parse(time.RFC3339, value)

	if e != nil {
		return value
	}

	return library.FormatCompact(t)
}
