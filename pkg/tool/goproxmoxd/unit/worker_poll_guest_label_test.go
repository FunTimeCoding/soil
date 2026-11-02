package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/worker_tester"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
)

func TestPollGuestLabelSet(t *testing.T) {
	c := worker_tester.PopulatedClient("pve")
	t.Cleanup(c.Close)
	w, y, _ := worker_tester.NewPollWorker(t, "pve", c)
	defer w.Close()
	w.Poll()

	if e := testutil.GatherAndCompare(
		y,
		strings.NewReader(worker_tester.GuestStatusExposition()),
		"proxmox_guest_status",
	); e != nil {
		t.Errorf("%v", e)
	}
}
