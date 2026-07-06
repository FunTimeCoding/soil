package power_panel

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestPanel(t *testing.T) {
	assert.NotNil(t, New(&netbox.PowerPanel{}))
}
