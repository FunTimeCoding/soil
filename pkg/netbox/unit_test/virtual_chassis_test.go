package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_chassis"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestChassis(t *testing.T) {
	assert.NotNil(t, virtual_chassis.New(&netbox.VirtualChassis{}))
}
