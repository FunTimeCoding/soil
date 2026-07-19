package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/rack_type"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestRackType(t *testing.T) {
	assert.NotNil(t, rack_type.New(&netbox.RackType{}))
}
