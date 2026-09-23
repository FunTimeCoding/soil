package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/worker_tester"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
)

func TestPollCollectsResources(t *testing.T) {
	w, y := worker_tester.NewPollWorker(
		"pve",
		worker_tester.PopulatedClient("pve"),
	)
	w.Poll()
	assert.Integer(t, 2, testutil.CollectAndCount(y, "proxmox_guest_status"))
	assert.Integer(t, 1, testutil.CollectAndCount(y, "proxmox_node_status"))
	assert.Integer(
		t,
		1,
		testutil.CollectAndCount(y, "proxmox_guest_backup_missing"),
	)
	assert.Integer(
		t,
		1,
		testutil.CollectAndCount(y, "proxmox_node_update_pending"),
	)

	if e := testutil.GatherAndCompare(
		y,
		strings.NewReader(worker_tester.ScrapeSuccess("1")),
		"proxmox_scrape_success",
	); e != nil {
		t.Errorf("%v", e)
	}
}

func TestPollFailureClearsResources(t *testing.T) {
	c := worker_tester.PopulatedClient("pve")
	w, y := worker_tester.NewPollWorker("pve", c)
	w.Poll()
	c.SetFailure(errors.New("hypervisor unreachable"))
	w.Poll()
	assert.Integer(t, 0, testutil.CollectAndCount(y, "proxmox_guest_status"))
	assert.Integer(t, 0, testutil.CollectAndCount(y, "proxmox_node_status"))

	if e := testutil.GatherAndCompare(
		y,
		strings.NewReader(worker_tester.ScrapeSuccess("0")),
		"proxmox_scrape_success",
	); e != nil {
		t.Errorf("%v", e)
	}
}
