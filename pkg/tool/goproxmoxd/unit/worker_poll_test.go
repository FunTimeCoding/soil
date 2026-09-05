package unit

import (
	"context"
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/log/logger"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
	"time"
)

func newPollWorker(
	name string,
	c *mock_client.Client,
) (*worker.Worker, *prometheus.Registry) {
	y := prometheus.NewRegistry()

	return worker.New(
		mock_service.New(name, c, nil),
		time.Minute,
		y,
		logger.New(context.Background()),
		reporter.NewOptional(constant.Identity.Name(), "test"),
	), y
}

func populatedClient(name string) *mock_client.Client {
	c := mock_client.New()
	c.AddNode(name)
	c.AddMachine(name, 100, "first")
	c.SetMachineStatus(name, 100, proxmoxConstant.RunningStatus)
	c.AddMachine(name, 101, "second")
	c.AddGuestNotInBackup(proxmoxConstant.MachineType, 101, "second")
	c.AddUpdatePending(name, "pve-docs", "9.2.3", "9.2.4")

	return c
}

func scrapeSuccess(value string) string {
	return strings.Join(
		[]string{
			"# HELP proxmox_scrape_success Whether the last poll of the hypervisor succeeded",
			"# TYPE proxmox_scrape_success gauge",
			fmt.Sprintf("proxmox_scrape_success{hypervisor=\"pve\"} %s", value),
			"",
		},
		"\n",
	)
}

func TestPollCollectsResources(t *testing.T) {
	w, y := newPollWorker("pve", populatedClient("pve"))
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
		strings.NewReader(scrapeSuccess("1")),
		"proxmox_scrape_success",
	); e != nil {
		t.Errorf("%v", e)
	}
}

func TestPollFailureClearsResources(t *testing.T) {
	c := populatedClient("pve")
	w, y := newPollWorker("pve", c)
	w.Poll()
	c.SetFailure(errors.New("hypervisor unreachable"))
	w.Poll()
	assert.Integer(t, 0, testutil.CollectAndCount(y, "proxmox_guest_status"))
	assert.Integer(t, 0, testutil.CollectAndCount(y, "proxmox_node_status"))

	if e := testutil.GatherAndCompare(
		y,
		strings.NewReader(scrapeSuccess("0")),
		"proxmox_scrape_success",
	); e != nil {
		t.Errorf("%v", e)
	}
}
