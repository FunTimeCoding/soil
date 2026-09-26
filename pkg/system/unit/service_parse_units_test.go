package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/service"
	"testing"
)

func TestParseUnitsReadsTheActiveColumn(t *testing.T) {
	v := service.ParseUnits(unitOutput())
	assert.String(t, "running", v["avahi-daemon.service"])
	assert.String(t, "failed", v["foxtrot.service"])
	assert.String(t, "stopped", v["fstrim.service"])
}

func TestParseUnitsRefusesAnUnloadedUnit(t *testing.T) {
	v := service.ParseUnits(unitOutput())
	assert.Integer(t, 4, len(v))
	_, okay := v["NetworkManager.service"]
	assert.False(t, okay)
}
