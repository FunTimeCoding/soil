package netboot

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"http://ftp.debian.org/debian/dists/Alfa/main/installer-Bravo/current/images/netboot/netboot.tar.gz",
		Link(upper.Alfa, upper.Bravo),
	)
}
