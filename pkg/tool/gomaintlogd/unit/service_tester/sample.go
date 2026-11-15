package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"

func Sample() *entry.Entry {
	v := entry.Stub()
	v.Action = "restart"
	v.User = "alice"

	return v
}
