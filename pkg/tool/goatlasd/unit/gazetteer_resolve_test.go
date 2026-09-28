package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"testing"
)

func TestResolveHostPrefersAddress(t *testing.T) {
	p, okay := gazetteer_tester.NewGazetteer().ResolveHost(
		"198.51.100.8",
		"02:aa:bb:cc:dd:08",
	)
	assert.True(t, okay)
	assert.String(t, "delta", p.Name)
	assert.String(t, "dcim.device", p.Kind)
}

func TestResolveHostFallsBackToHardware(t *testing.T) {
	p, okay := gazetteer_tester.NewGazetteer().ResolveHost(
		"198.51.100.4",
		"02:aa:bb:cc:dd:04",
	)
	assert.True(t, okay)
	assert.String(t, "echo", p.Name)
	assert.String(t, "virtualization.virtualmachine", p.Kind)
}

func TestResolveHostLeavesUnknownHostsUnclaimed(t *testing.T) {
	_, okay := gazetteer_tester.NewGazetteer().ResolveHost(
		"198.51.100.233",
		"02:aa:bb:cc:dd:33",
	)
	assert.False(t, okay)
}

func TestResolveHostIgnoresHardwareAddressCase(t *testing.T) {
	p, okay := gazetteer_tester.NewGazetteer().ResolveHost(
		"198.51.100.4",
		"02:AA:BB:CC:DD:04",
	)
	assert.True(t, okay)
	assert.String(t, "echo", p.Name)
}
