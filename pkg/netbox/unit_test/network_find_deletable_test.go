package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"testing"
)

func TestFindDeletable(t *testing.T) {
	assert.Any(
		t,
		[]*network.Interface{{Name: "eth1"}},
		network.FindDeletable(
			[]*network.Interface{{Name: network.Eth1}},
			[]*network.Definition{{Name: network.Eth0}},
		),
	)
}
