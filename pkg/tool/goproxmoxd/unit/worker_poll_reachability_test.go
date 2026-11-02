package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/worker_tester"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
)

func TestARefusedHypervisorIsLeftToTheGuard(t *testing.T) {
	c := worker_tester.PopulatedClient("pve")
	defer c.Close()
	w, y, r := worker_tester.NewPollWorker(t, "pve", c)
	defer w.Close()
	w.Poll()
	c.SetFailure(
		connection.New(
			errorsConstant.Unreachable,
			"pve",
			"",
			errorsConstant.Refused,
		),
	)
	w.Poll()
	c.SetFailure(nil)
	w.Poll()
	assert.Count(t, 0, r.Events())

	if e := testutil.GatherAndCompare(
		y,
		strings.NewReader(worker_tester.ScrapeSuccess("0")),
		"proxmox_scrape_success",
	); e != nil {
		t.Errorf("%v", e)
	}
}

func TestAPollFailingOtherwiseIsCaptured(t *testing.T) {
	c := worker_tester.PopulatedClient("pve")
	defer c.Close()
	w, _, r := worker_tester.NewPollWorker(t, "pve", c)
	defer w.Close()
	w.Poll()
	c.SetFailure(errors.New("token rejected"))
	w.Poll()
	assert.Count(t, 1, r.Events())
}
