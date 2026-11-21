package upgrade_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"

func New() *Result {
	return &Result{Groups: []*upgrade_group.Group{}, Releases: []string{}}
}
