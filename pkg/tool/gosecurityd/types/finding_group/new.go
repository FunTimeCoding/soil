package finding_group

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"

func New(
	vulnerability string,
	score float64,
	severity string,
) *Group {
	return &Group{
		Vulnerability: vulnerability,
		Score:         score,
		Severity:      severity,
		Unimportant:   true,
		Count:         affected_count.New(),
	}
}
