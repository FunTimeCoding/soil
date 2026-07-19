package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_template"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestConfigurationTemplate(t *testing.T) {
	assert.NotNil(t, configuration_template.New(&netbox.ConfigTemplate{}))
}
