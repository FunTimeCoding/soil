package system_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/kernel_offer"
	"time"
)

type Summary struct {
	Name            string
	Release         string
	Kernel          string
	RunningKernel   string
	NewestKernel    string
	KernelState     string
	KernelReason    string
	KernelFindings  int
	KernelWithFix   int
	KernelOffers    []*kernel_offer.Offer
	ListsUpdated    *time.Time
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
	Reboot          int
	Critical        int
	TopScore        float64
	Updated         time.Time
	TrackerVerified time.Time
}
