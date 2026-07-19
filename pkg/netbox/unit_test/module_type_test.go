package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/module_type"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestModuleType(t *testing.T) {
	assert.NotNil(t, module_type.New(&netbox.ModuleType{}))
}
