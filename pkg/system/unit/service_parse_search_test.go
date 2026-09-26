package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/service"
	"testing"
)

func TestParseSearchMapsPathToPackage(t *testing.T) {
	v := service.ParseSearch(searchOutput())
	assert.String(
		t,
		"avahi-daemon",
		v["/usr/lib/systemd/system/avahi-daemon.service"],
	)
	assert.String(t, "foxtrot", v["/lib/systemd/system/foxtrot.service"])
}

func TestParseSearchRefusesTheNotFoundLine(t *testing.T) {
	assert.Integer(t, 2, len(service.ParseSearch(searchOutput())))
}
