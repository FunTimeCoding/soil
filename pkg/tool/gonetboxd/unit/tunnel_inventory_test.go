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

func TestInventoryDialsHostsOffSiteAtTunnelAddress(t *testing.T) {
	c := service_tester.NewSource(t)
	seedTunnelHosts(c)
	i, e := service_tester.Service(t, c).AnsibleInventory(
		constant.AnsibleField,
		"kilo",
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                alfa:
                    ansible_host: 198.51.100.1
                bravo:
                    ansible_host: 192.0.2.6
                charlie:
                    ansible_host: 198.51.100.9
`,
		i.Markup(),
	)
}

func TestInventoryWithoutSiteKeepsPrimaryAddresses(t *testing.T) {
	c := service_tester.NewSource(t)
	seedTunnelHosts(c)
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
                alfa:
                    ansible_host: 203.0.113.5
                bravo:
                    ansible_host: 192.0.2.6
                charlie: {}
`,
		i.Markup(),
	)
}

func TestInventoryOffSiteWithoutTunnelKeepsPrimaryAddress(t *testing.T) {
	c := service_tester.NewSource(t)
	delta := c.SeedAddressedDevice("delta", "203.0.113.7/32")
	delta.Site = "lima"
	delta.CustomFields = map[string]any{constant.AnsibleField: []string{"echo"}}
	c.SeedHostAddress("delta", "198.51.100.2/24", "dcim.interface", 9)
	c.SeedTunnelTermination("dcim.interface", 10)
	i, e := service_tester.Service(t, c).AnsibleInventory(
		constant.AnsibleField,
		"kilo",
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                delta:
                    ansible_host: 203.0.113.7
`,
		i.Markup(),
	)
}

func TestInventoryRouteTakesCallerSite(t *testing.T) {
	c := service_tester.NewSource(t)
	seedTunnelHosts(c)
	d := daemon.New(
		service_tester.NewOption(t, c),
		logger.New(t.Context()),
		memory.New(),
		metric.New(),
	)
	t.Cleanup(d.Close)
	site := "kilo"
	response, e := newClient(
		t,
		newDaemonServer(t, d, memory.New()),
	).GetAnsibleInventoryWithResponse(
		context.Background(),
		&client.GetAnsibleInventoryParams{Site: &site},
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                alfa:
                    ansible_host: 198.51.100.1
                bravo:
                    ansible_host: 192.0.2.6
                charlie:
                    ansible_host: 198.51.100.9
`,
		string(response.Body),
	)
}
