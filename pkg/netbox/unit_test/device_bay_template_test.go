package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/device_bay_template"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestDeviceBayTemplate(t *testing.T) {
	assert.NotNil(t, device_bay_template.New(&netbox.DeviceBayTemplate{}))
}
