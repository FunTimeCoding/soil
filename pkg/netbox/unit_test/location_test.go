package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestLocation(t *testing.T) {
	assert.NotNil(t, location.New(&netbox.Location{}))
}
