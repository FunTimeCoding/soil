package virtual_chassis

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestChassis(t *testing.T) {
	assert.NotNil(t, New(&netbox.VirtualChassis{}))
}
