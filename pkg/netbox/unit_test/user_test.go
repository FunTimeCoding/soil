package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/user"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestUser(t *testing.T) {
	assert.NotNil(t, user.New(&netbox.User{}))
}
