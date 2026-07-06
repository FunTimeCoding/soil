package module_type_profile

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestProfile(t *testing.T) {
	assert.NotNil(t, New(&netbox.ModuleTypeProfile{}))
}
