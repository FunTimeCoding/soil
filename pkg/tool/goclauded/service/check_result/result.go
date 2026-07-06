package check_result

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

type Result struct {
	Callsign string
	Changed  bool
	Entries  []queue.Entry
}
