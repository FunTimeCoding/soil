package state_tree

import "github.com/funtimecoding/soil/pkg/salt/types/top_entry"

func isLive(e *top_entry.Entry) bool {
	return e.Live
}
