package configuration_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestContext(t *testing.T) {
	assert.NotNil(t, New(&netbox.ConfigContext{}))
}
