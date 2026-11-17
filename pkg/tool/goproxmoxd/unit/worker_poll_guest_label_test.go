package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/worker_tester"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
)

func TestPollGuestLabelSet(t *testing.T) {
	o := worker_tester.New(t)
	o.Populate()
	w := o.Worker()
	w.Poll()

	if e := testutil.GatherAndCompare(
		o.Registry,
		strings.NewReader(worker_tester.GuestStatusExposition()),
		"proxmox_guest_status",
	); e != nil {
		t.Errorf("%v", e)
	}
}
