package matched

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"
	"time"
)

type Matched struct {
	Updated       time.Time
	Trackers      []*tracker.Tracker
	Scores        uint64
	Findings      []*system_result.Finding
	RunningKernel string
	NewestKernel  string
	KernelChecked bool
}
