package retarget_report

import "github.com/funtimecoding/soil/pkg/lint/types/retarget_change"

type Report struct {
	Identifier int64
	Name       string
	Changes    []*retarget_change.Change
}
