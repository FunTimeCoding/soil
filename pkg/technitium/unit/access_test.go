package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/policy"
	"testing"
)

func TestWidenAccessOpensDenyToTheNetworkList(t *testing.T) {
	assert.String(
		t,
		"UseSpecifiedNetworkACL",
		technitium.WidenAccess(constant.AccessDeny),
	)
}

func TestWidenAccessKeepsZoneNameServersAccess(t *testing.T) {
	assert.String(
		t,
		"AllowZoneNameServersAndUseSpecifiedNetworkACL",
		technitium.WidenAccess(constant.AccessNameServer),
	)
}

func TestWidenAccessNeverNarrowsAnOpenZone(t *testing.T) {
	assert.String(t, "Allow", technitium.WidenAccess(constant.AccessAllow))
}

func TestAllowUpdateOnAFreshZone(t *testing.T) {
	o := zoneOptions(
		`{"name":"alfa.example","update":"Deny","updateNetworkACL":[],"updateSecurityPolicies":[]}`,
	)
	technitium.AllowUpdate(
		o,
		"198.51.100.0/24",
		policy.New("alfa", "*.alfa.example", []string{"A", "TXT"}),
	)
	assert.String(t, "UseSpecifiedNetworkACL", o.Update)
	assert.Strings(t, []string{"198.51.100.0/24"}, o.UpdateNetwork)
	assert.Count(t, 1, o.UpdatePolicies)
	assert.String(t, "*.alfa.example", o.UpdatePolicies[0].Domain)
}

func TestAllowUpdateTwiceAddsNothingTwice(t *testing.T) {
	o := zoneOptions(`{"name":"alfa.example","update":"Deny"}`)
	p := policy.New("alfa", "*.alfa.example", []string{"A"})
	technitium.AllowUpdate(o, "198.51.100.0/24", p)
	technitium.AllowUpdate(o, "198.51.100.0/24", p)
	assert.Count(t, 1, o.UpdateNetwork)
	assert.Count(t, 1, o.UpdatePolicies)
}

func TestAllowUpdateKeepsExistingNetworksAndPolicies(t *testing.T) {
	o := zoneOptions(
		`{"name":"alfa.example","update":"UseSpecifiedNetworkACL",
		"updateNetworkACL":["192.0.2.0/24"],
		"updateSecurityPolicies":[{"tsigKeyName":"bravo","domain":"alfa.example","allowedTypes":["A"]}]}`,
	)
	technitium.AllowUpdate(
		o,
		"198.51.100.0/24",
		policy.New("alfa", "*.alfa.example", []string{"A"}),
	)
	assert.Strings(
		t,
		[]string{"192.0.2.0/24", "198.51.100.0/24"},
		o.UpdateNetwork,
	)
	assert.Count(t, 2, o.UpdatePolicies)
}

func TestAllowTransferKeepsZoneNameServersAndAddsTheKey(t *testing.T) {
	o := zoneOptions(
		`{"name":"alfa.example","zoneTransfer":"AllowOnlyZoneNameServers",
		"zoneTransferNetworkACL":[],"zoneTransferTsigKeyNames":[]}`,
	)
	technitium.AllowTransfer(o, "198.51.100.0/24", "alfa")
	assert.String(
		t,
		"AllowZoneNameServersAndUseSpecifiedNetworkACL",
		o.ZoneTransfer,
	)
	assert.Strings(t, []string{"198.51.100.0/24"}, o.ZoneTransferNetwork)
	assert.Strings(t, []string{"alfa"}, o.ZoneTransferKeyNames)
}
