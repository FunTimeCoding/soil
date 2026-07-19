package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/module"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestModule(t *testing.T) {
	assert.NotNil(t, module.New(&netbox.Module{}))
}
