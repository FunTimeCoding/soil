package system_result

import "time"

type Summary struct {
	Name            string
	Release         string
	Kernel          string
	Priority        int
	Labels          map[string]string
	Coverage        string
	Checked         bool
	Clean           bool
	Stale           bool
	TrackerStale    bool
	Findings        int
	NoFix           int
	PaidOnly        int
	Upgradable      int
	Updated         time.Time
	TrackerVerified time.Time
}
