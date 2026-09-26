package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/service"
	"testing"
)

func TestParsePolicyTakesTheInstalledVersionSource(t *testing.T) {
	v := service.ParsePolicy(policyOutput())
	assert.String(t, "1.2.3", v["foxtrot"].Version)
	assert.String(t, "https://apt.example.org", v["foxtrot"].Origin)
}

func TestParsePolicySeparatesPackages(t *testing.T) {
	v := service.ParsePolicy(policyOutput())
	assert.Integer(t, 2, len(v))
	assert.String(t, "0.8-16", v["avahi-daemon"].Version)
	assert.String(
		t,
		"mirror+file:/etc/apt/mirrors/debian.list",
		v["avahi-daemon"].Origin,
	)
}
