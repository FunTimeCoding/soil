package trace_report

import "github.com/funtimecoding/soil/pkg/salt/types/top_entry"

type Report struct {
	Target  string             `json:"target"`
	Summary *Summary           `json:"summary"`
	Entries []*top_entry.Entry `json:"entries"`
	States  []*Node            `json:"states"`
	Files   []*Node            `json:"files"`
	Missing []string           `json:"missing"`
}
