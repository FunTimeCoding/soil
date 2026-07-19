package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/contact"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestContact(t *testing.T) {
	assert.NotNil(t, contact.New(&netbox.Contact{}))
}
