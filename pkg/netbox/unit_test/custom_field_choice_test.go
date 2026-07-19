package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field_choice"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestChoice(t *testing.T) {
	assert.NotNil(t, custom_field_choice.New(&netbox.CustomFieldChoiceSet{}))
}
