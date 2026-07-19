package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestTenant(t *testing.T) {
	assert.NotNil(t, tenant.New(&netbox.Tenant{}))
}
