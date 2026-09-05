package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"testing"
)

func TestFindDeletable(t *testing.T) {
	assert.Any(
		t,
		[]*network.Interface{{Name: "eth1"}},
		network.FindDeletable(
			[]*network.Interface{{Name: constant.Eth1}},
			[]*network.Definition{{Name: constant.Eth0}},
		),
	)
}
