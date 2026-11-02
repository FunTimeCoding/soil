package retarget_report

import "github.com/funtimecoding/soil/pkg/lint/types/retarget_change"

func New(
	identifier int64,
	name string,
	changes []*retarget_change.Change,
) *Report {
	return &Report{Identifier: identifier, Name: name, Changes: changes}
}
