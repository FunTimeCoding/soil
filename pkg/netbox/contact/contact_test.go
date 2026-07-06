package contact

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestContact(t *testing.T) {
	assert.NotNil(t, New(&netbox.Contact{}))
}
