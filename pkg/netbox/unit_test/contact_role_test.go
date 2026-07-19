package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/contact_role"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestContactRole(t *testing.T) {
	assert.NotNil(t, contact_role.New(&netbox.ContactRole{}))
}
