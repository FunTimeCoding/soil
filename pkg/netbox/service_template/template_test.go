package service_template

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestTemplate(t *testing.T) {
	assert.NotNil(t, New(&netbox.ServiceTemplate{}))
}
