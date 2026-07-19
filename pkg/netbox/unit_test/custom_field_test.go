package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestField(t *testing.T) {
	assert.NotNil(t, custom_field.New(&netbox.CustomField{}))
}
