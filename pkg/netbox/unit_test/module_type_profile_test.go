package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/module_type_profile"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestProfile(t *testing.T) {
	assert.NotNil(t, module_type_profile.New(&netbox.ModuleTypeProfile{}))
}
