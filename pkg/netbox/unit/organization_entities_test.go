package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/contact"
	"github.com/funtimecoding/soil/pkg/netbox/contact_group"
	"github.com/funtimecoding/soil/pkg/netbox/contact_role"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/netbox/tenant_group"
	"github.com/funtimecoding/soil/pkg/netbox/user"
	"github.com/funtimecoding/soil/pkg/netbox/user_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestOrganizationEntities(t *testing.T) {
	assert.NotNil(t, contact_group.New(&netbox.ContactGroup{}))
	assert.NotNil(t, contact_role.New(&netbox.ContactRole{}))
	assert.NotNil(t, contact.New(&netbox.Contact{}))
	assert.NotNil(t, tenant_group.New(&netbox.TenantGroup{}))
	assert.NotNil(t, tenant.New(&netbox.Tenant{}))
	assert.NotNil(t, user_group.New(&netbox.Group{}))
	assert.NotNil(t, user.New(&netbox.User{}))
}
