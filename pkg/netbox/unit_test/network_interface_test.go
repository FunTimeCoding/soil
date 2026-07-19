package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestInterface(t *testing.T) {
	assert.NotNil(t, network.New(&netbox.Interface{}))
}
