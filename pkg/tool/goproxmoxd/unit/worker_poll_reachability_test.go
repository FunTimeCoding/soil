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

func TestRefusedHypervisorIsLeftToGuard(t *testing.T) {
	o := worker_tester.New(t)
	o.Populate()
	w := o.Worker()
	w.Poll()
	o.Alfa().Hypervisor.SetFailure(
		connection.New(
			errorsConstant.Unreachable,
			"alfa",
			"",
			errorsConstant.Refused,
		),
	)
	w.Poll()
	o.Alfa().Hypervisor.SetFailure(nil)
	w.Poll()
	assert.Count(t, 0, o.Reporter.Events())

	if e := testutil.GatherAndCompare(
		o.Registry,
		strings.NewReader(worker_tester.ScrapeSuccess("0")),
		"proxmox_scrape_success",
	); e != nil {
		t.Errorf("%v", e)
	}
}

func TestPollFailingOtherwiseIsCaptured(t *testing.T) {
	o := worker_tester.New(t)
	o.Populate()
	w := o.Worker()
	w.Poll()
	o.Alfa().Hypervisor.SetFailure(errors.New("token rejected"))
	w.Poll()
	assert.Count(t, 1, o.Reporter.Events())
}
