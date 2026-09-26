package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/service"
	"testing"
)

func TestOwnerOfFindsTheUnmergedSpelling(t *testing.T) {
	assert.String(
		t,
		"foxtrot",
		service.OwnerOf(
			map[string]string{"/lib/systemd/system/foxtrot.service": "foxtrot"},
			"/usr/lib/systemd/system/foxtrot.service",
		),
	)
}

func TestOwnerOfFindsTheMergedSpelling(t *testing.T) {
	assert.String(
		t,
		"avahi-daemon",
		service.OwnerOf(
			map[string]string{
				"/usr/lib/systemd/system/avahi-daemon.service": "avahi-daemon",
			},
			"/lib/systemd/system/avahi-daemon.service",
		),
	)
}

func TestOwnerOfRefusesAnUnownedPath(t *testing.T) {
	assert.String(
		t,
		"",
		service.OwnerOf(
			map[string]string{},
			"/etc/systemd/system/custom.service",
		),
	)
}
