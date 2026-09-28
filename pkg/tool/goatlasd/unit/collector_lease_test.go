package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/mock_lease_source"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"testing"
)

func TestLeaseResolvesByAddress(t *testing.T) {
	f := mock_lease_source.New()
	f.Add("delta", "198.51.100.8", "02:aa:bb:cc:dd:08", true)
	v := collector_tester.CollectLeases(t, f)
	assert.Count(t, 1, v)
	assert.String(t, "delta", v[0].PlaceName)
	assert.String(t, "dcim.device", v[0].PlaceKind)
}

func TestLeaseResolvesByHardwareAddressWhenAddressIsUnknown(t *testing.T) {
	f := mock_lease_source.New()
	f.Add(
		constant.FixtureVirtualMachineNode,
		"198.51.100.4",
		"02:aa:bb:cc:dd:04",
		true,
	)
	v := collector_tester.CollectLeases(t, f)
	assert.String(t, "echo", v[0].PlaceName)
	assert.String(t, "virtualization.virtualmachine", v[0].PlaceKind)
}

func TestLeaseNeverResolvesByHostname(t *testing.T) {
	f := mock_lease_source.New()
	f.Add("delta", "198.51.100.99", "aa:bb:cc:dd:ee:ff", false)
	v := collector_tester.CollectLeases(t, f)
	assert.String(t, "", v[0].PlaceKind)
	assert.String(t, "delta", v[0].Hostname)
}

func TestLeaseKeepsUnclaimedHosts(t *testing.T) {
	f := mock_lease_source.New()
	f.Add("juliet", "198.51.100.233", "02:aa:bb:cc:dd:33", false)
	v := collector_tester.CollectLeases(t, f)
	assert.Count(t, 1, v)
	assert.String(t, "", v[0].PlaceKind)
	assert.String(t, "02:aa:bb:cc:dd:33", v[0].HardwareAddress)
	assert.False(t, v[0].Reserved)
}

func TestLeaseCarriesTheReservationFlag(t *testing.T) {
	f := mock_lease_source.New()
	f.Add("delta", "198.51.100.8", "02:aa:bb:cc:dd:08", true)
	v := collector_tester.CollectLeases(t, f)
	assert.True(t, v[0].Reserved)
	assert.String(t, "gopnsensed", v[0].Source)
}
