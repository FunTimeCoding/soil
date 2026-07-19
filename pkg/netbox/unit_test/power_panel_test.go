package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/power_panel"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestPanel(t *testing.T) {
	assert.NotNil(t, power_panel.New(&netbox.PowerPanel{}))
}
