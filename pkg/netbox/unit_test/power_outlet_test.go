package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestOutlet(t *testing.T) {
	assert.NotNil(t, power_outlet.New(&netbox.PowerOutlet{}))
}
