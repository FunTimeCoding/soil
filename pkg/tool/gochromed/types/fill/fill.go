package fill

import "github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"

type Fill struct {
	Selector *selector.Selector
	UID      string
	Value    string
	Direct   bool
	Snapshot bool
}
