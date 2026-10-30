package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"testing"
)

func TestTheMockReturnsEachMachinesOwnDisks(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	alfa := m.SeedSizedVirtualMachine("alfa", 4, 8192, "charlie")
	bravo := m.SeedVirtualMachine("bravo")
	m.SeedVirtualDisk(alfa, "root", 40960, map[string]any{"tier": "delta"})
	m.SeedVirtualDisk(alfa, "data", 102400, nil)
	m.SeedVirtualDisk(bravo, "root", 20480, nil)
	disks, e := m.VirtualMachineDisks(alfa)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, disks)
	assert.String(t, "root", disks[0].Name)
	assert.Integer(t, 40960, int(disks[0].Size))
	assert.Any(t, "delta", disks[0].Raw.CustomFields["tier"])
	assert.String(t, "alfa", disks[0].Raw.VirtualMachine.GetName())
	others, f := m.VirtualMachineDisks(bravo)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, others)
}

func TestASizedMachineCarriesItsCoresMemoryAndSite(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	m.SeedSizedVirtualMachine("alfa", 4, 8192, "charlie")
	v, e := m.VirtualMachineByName("alfa")
	assert.FatalOnError(t, e)
	assert.String(t, "charlie", v.Site)
	assert.Any(t, 4.0, v.Raw.GetVcpus())
	assert.Integer(t, 8192, int(v.Raw.GetMemory()))
}

func TestACreatedDiskIsReturnedWithItsMachine(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	alfa := m.SeedVirtualMachine("alfa")
	_, e := m.CreateVirtualDisk("alfa", "swap", 4096)
	assert.FatalOnError(t, e)
	disks, f := m.VirtualMachineDisks(alfa)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, disks)
	assert.String(t, "swap", disks[0].Name)
	_, g := m.CreateVirtualDisk("charlie", "swap", 4096)
	assert.True(t, not_found.Is(g))
}

func TestCoresAndMemoryAreSetOnTheMachine(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	m.SeedSizedVirtualMachine("alfa", 2, 4096, "echo")
	_, e := m.SetVirtualMachineCores("alfa", 8)
	assert.FatalOnError(t, e)
	_, f := m.SetVirtualMachineMemory("alfa", 16384)
	assert.FatalOnError(t, f)
	v, g := m.VirtualMachineByName("alfa")
	assert.FatalOnError(t, g)
	assert.Any(t, 8.0, v.Raw.GetVcpus())
	assert.Integer(t, 16384, int(v.Raw.GetMemory()))
	_, h := m.SetVirtualMachineCores("charlie", 8)
	assert.True(t, not_found.Is(h))
}

func TestTheDisksFailOnTheirOwn(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	alfa := m.SeedVirtualMachine("alfa")
	m.Fail(m.VirtualMachineDisks, unavailable.Format("netbox: down"))
	_, e := m.VirtualMachineDisks(alfa)
	assert.True(t, unavailable.Is(e))
	_, f := m.VirtualMachineByName("alfa")
	assert.FatalOnError(t, f)
}
