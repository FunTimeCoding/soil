package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestRack(t *testing.T) {
	assert.NotNil(t, rack.New(&netbox.Rack{}))
}
