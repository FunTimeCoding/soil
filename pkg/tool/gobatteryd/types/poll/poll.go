package poll

import "github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/reading"

type Poll struct {
	Readings  []*reading.Reading
	Reachable map[string]bool
	Faults    []error
}
