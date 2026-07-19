package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestBay(t *testing.T) {
	assert.NotNil(t, module_bay.New(&netbox.ModuleBay{}))
}
