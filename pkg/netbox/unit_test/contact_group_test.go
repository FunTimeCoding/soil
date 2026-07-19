package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/contact_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestContactGroup(t *testing.T) {
	assert.NotNil(t, contact_group.New(&netbox.ContactGroup{}))
}
