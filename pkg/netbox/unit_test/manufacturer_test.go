package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestManufacturer(t *testing.T) {
	assert.NotNil(t, manufacturer.New(&netbox.Manufacturer{}))
}
