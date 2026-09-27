package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/service"
	"testing"
)

func TestNewUnitServiceCarriesTheOwningPackage(t *testing.T) {
	v := service.NewUnitService(
		"foxtrot.service",
		constant.ServiceRunning,
		"/lib/systemd/system/foxtrot.service",
		"foxtrot",
		servicePolicies(),
		map[string]bool{"foxtrot": true},
	)
	assert.String(t, "foxtrot.service", v.Name)
	assert.String(t, "foxtrot", v.Package)
	assert.String(t, "1.2.3", v.Version)
	assert.True(t, v.Deliberate)
}

func TestNewUnitServiceWithoutPackageCarriesNone(t *testing.T) {
	v := service.NewUnitService(
		"charlie.service",
		constant.ServiceFailed,
		"/etc/systemd/system/charlie.service",
		"",
		servicePolicies(),
		map[string]bool{},
	)
	assert.String(t, "charlie.service", v.Name)
	assert.String(t, "", v.Package)
	assert.String(t, "", v.Version)
}
