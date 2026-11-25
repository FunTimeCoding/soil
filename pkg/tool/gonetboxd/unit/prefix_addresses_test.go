package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	gonetboxdMock "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/client_tester"
	"testing"
)

func TestPrefixAddressesNameWhatHoldsEach(t *testing.T) {
	m := mock_client.New(t)
	vip := m.SeedHostAddress("alfa", "192.0.2.5/32", "dcim.interface", 1)
	vip.Role = "vip"
	vip.DomainName = "bravo.example.org"
	m.SeedHostAddress(
		"charlie",
		"192.0.2.9/32",
		"virtualization.vminterface",
		2,
	)
	m.SeedHostAddress("delta", "198.51.100.1/32", "dcim.interface", 3)
	v, e := client_tester.New(t, m).ListPrefixAddresses(
		context.Background(),
		"192.0.2.0/24",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
	assert.String(t, "192.0.2.5", v[0].Address)
	assert.String(t, "alfa", *v[0].Device)
	assert.String(t, "vip", *v[0].Role)
	assert.String(t, "bravo.example.org", *v[0].DomainName)
	assert.True(t, v[0].VirtualMachine == nil)
	assert.String(t, "charlie", *v[1].VirtualMachine)
	assert.True(t, v[1].Device == nil)
}

func TestPrefixAddressesRefuseNonCidrBeforeNetBox(t *testing.T) {
	m := mock_client.New(t)
	m.Failure = unavailable.Format("netbox down")
	_, e := client_tester.New(t, m).ListPrefixAddresses(
		context.Background(),
		"192.0.2.0",
	)
	assert.True(t, validation.Is(e))
}

func TestPrefixAddressesToolNamesHolder(t *testing.T) {
	s := base.New(t)
	s.Source.SeedHostAddress("alfa", "192.0.2.5/32", "dcim.interface", 1)
	tools := s.Tools()
	assert.StringContains(
		t,
		`"device": "alfa"`,
		tools.MustCallTool(
			constant.ListPrefixAddresses,
			map[string]any{constant.Prefix: "192.0.2.0/24"},
		),
	)
	assert.StringContains(
		t,
		"not in CIDR notation",
		tools.MustCallToolError(
			constant.ListPrefixAddresses,
			map[string]any{constant.Prefix: "alfa"},
		),
	)
}

func TestDaemonMockListsPrefixLikeDaemon(t *testing.T) {
	m := mock_client.New(t)
	m.SeedHostAddress("alfa", "192.0.2.5/32", "dcim.interface", 1)
	m.SeedHostAddress("bravo", "192.0.2.130/32", "dcim.interface", 2)
	daemon := client_tester.New(t, m)
	every, e := daemon.ListPrefixAddresses(context.Background(), "192.0.2.0/24")
	assert.FatalOnError(t, e)
	mock := gonetboxdMock.New(t)
	mock.AddressList = every

	for _, c := range []face.Client{daemon, mock} {
		v, f := c.ListPrefixAddresses(context.Background(), "192.0.2.128/25")
		assert.FatalOnError(t, f)
		assert.Count(t, 1, v)
		assert.String(t, "bravo", *v[0].Device)
		_, g := c.ListPrefixAddresses(context.Background(), "bravo")
		assert.True(t, validation.Is(g))
	}
}
