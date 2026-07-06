package network

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFindByName(t *testing.T) {
	interfaces := []*Interface{{Name: Eth0, Type: InterfaceVirtual}}
	// Happy path
	assert.Any(
		t,
		&Interface{Name: "eth0", Type: "virtual"},
		FindByName(interfaces, Eth0),
	)
	// Not found
	var expected *Interface
	assert.Any(t, expected, FindByName(interfaces, Eth1))
}
