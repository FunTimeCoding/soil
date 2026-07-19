package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/front_port"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestFrontPort(t *testing.T) {
	assert.NotNil(t, front_port.New(&netbox.FrontPort{}))
}
