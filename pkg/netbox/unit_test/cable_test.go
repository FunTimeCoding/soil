package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestCable(t *testing.T) {
	assert.NotNil(t, cable.New(&netbox.Cable{}))
}
