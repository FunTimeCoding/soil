package unit_test

import (
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
)

func guestStatusExposition() string {
	return strings.Join(
		[]string{
			"# HELP proxmox_guest_status Guest status, one series per observed status with value 1",
			"# TYPE proxmox_guest_status gauge",
			`proxmox_guest_status{hypervisor="pve",identifier="100",name="first",node="pve",status="running",type="qemu"} 1`,
			`proxmox_guest_status{hypervisor="pve",identifier="101",name="second",node="pve",status="",type="qemu"} 1`,
			"",
		},
		"\n",
	)
}

func TestPollGuestLabelSet(t *testing.T) {
	w, y := newPollWorker("pve", populatedClient("pve"))
	w.Poll()

	if e := testutil.GatherAndCompare(
		y,
		strings.NewReader(guestStatusExposition()),
		"proxmox_guest_status",
	); e != nil {
		t.Errorf("%v", e)
	}
}
