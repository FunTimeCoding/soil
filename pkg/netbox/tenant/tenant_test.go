package tenant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestTenant(t *testing.T) {
	assert.NotNil(t, New(&netbox.Tenant{}))
}
