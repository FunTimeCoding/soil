package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/install/netboot"
	"testing"
)

func TestDebianNetbootLink(t *testing.T) {
	assert.String(
		t,
		"http://ftp.debian.org/debian/dists/Alfa/main/installer-Bravo/current/images/netboot/netboot.tar.gz",
		netboot.Link(stringsConstant.UpperAlfa, stringsConstant.UpperBravo),
	)
}
