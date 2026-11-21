package finding_group

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"

type Group struct {
	Vulnerability string
	Score         float64
	Severity      string
	Unimportant   bool
	Systems       []string
	Count         *affected_count.Count
}
