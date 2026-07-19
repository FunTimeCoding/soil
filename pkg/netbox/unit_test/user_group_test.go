package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/user_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestUserGroup(t *testing.T) {
	assert.NotNil(t, user_group.New(&netbox.Group{}))
}
