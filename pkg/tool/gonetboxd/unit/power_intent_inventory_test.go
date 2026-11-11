package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func TestPowerIntentServiceKeepsActiveHosts(t *testing.T) {
	c := service_tester.NewSource(t)
	seedIntentHosts(c)
	i, e := service_tester.Service(t, c).
		WithPowerIntent().
		AnsibleInventory(constant.AnsibleField, "")
	assert.FatalOnError(t, e)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                alfa: {}
                charlie: {}
`,
		i.Markup(),
	)
	machines, e := c.VirtualMachines()
	assert.FatalOnError(t, e)
	assert.Count(t, 2, machines)
	assert.String(t, "bravo", machines[1].Name)
	devices, e := c.Devices()
	assert.FatalOnError(t, e)
	assert.Count(t, 2, devices)
	assert.String(t, "delta", devices[1].Name)
}

func TestPowerIntentDaemonKeepsActiveHosts(t *testing.T) {
	c := service_tester.NewSource(t)
	seedIntentHosts(c)
	o := service_tester.NewOption(t, c)
	o.PowerIntent = true
	d := daemon.New(o, logger.New(t.Context()), memory.New(), metric.New())
	t.Cleanup(d.Close)
	response, e := newClient(
		t,
		newDaemonServer(t, d, memory.New()),
	).GetAnsibleInventoryWithResponse(
		context.Background(),
		&client.GetAnsibleInventoryParams{},
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                alfa: {}
                charlie: {}
`,
		string(response.Body),
	)
}

func TestInventoryKeepsEveryHost(t *testing.T) {
	c := service_tester.NewSource(t)
	seedIntentHosts(c)
	i, e := service_tester.Service(t, c).AnsibleInventory(
		constant.AnsibleField,
		"",
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                alfa: {}
                bravo: {}
                charlie: {}
                delta: {}
`,
		i.Markup(),
	)
}
