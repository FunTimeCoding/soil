package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/tenant_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestTenantGroup(t *testing.T) {
	assert.NotNil(t, tenant_group.New(&netbox.TenantGroup{}))
}
