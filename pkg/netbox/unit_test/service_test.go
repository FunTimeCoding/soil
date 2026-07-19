package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/service"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestService(t *testing.T) {
	assert.NotNil(t, service.New(&netbox.Service{}))
}
