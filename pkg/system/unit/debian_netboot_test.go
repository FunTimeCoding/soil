package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/debian/netboot"
	"testing"
)

func TestDebianNetbootLink(t *testing.T) {
	assert.String(
		t,
		"http://ftp.debian.org/debian/dists/Alfa/main/installer-Bravo/current/images/netboot/netboot.tar.gz",
		netboot.Link(constant.UpperAlfa, constant.UpperBravo),
	)
}
