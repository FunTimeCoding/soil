package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestConfigurationContext(t *testing.T) {
	assert.NotNil(t, configuration_context.New(&netbox.ConfigContext{}))
}
