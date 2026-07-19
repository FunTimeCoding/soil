package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/rack_role"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestRackRole(t *testing.T) {
	assert.NotNil(t, rack_role.New(&netbox.RackRole{}))
}
