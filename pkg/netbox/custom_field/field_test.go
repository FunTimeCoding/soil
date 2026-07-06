package custom_field

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestField(t *testing.T) {
	assert.NotNil(t, New(&netbox.CustomField{}))
}
