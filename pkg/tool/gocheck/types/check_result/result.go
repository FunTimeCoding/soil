package check_result

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/types/layer_change"
)

type Result struct {
	Root         string
	Announcement string
	Stale        string
	Unchanged    bool
	Sections     []*output.Results
	Changes      []*layer_change.Change
	Checked      []string
}
