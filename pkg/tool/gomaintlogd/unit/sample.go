package unit

import "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"

func sample() *entry.Entry {
	v := entry.New()
	v.Action = "restart"
	v.User = "alice"

	return v
}
