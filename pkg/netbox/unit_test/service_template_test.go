package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/service_template"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestServiceTemplate(t *testing.T) {
	assert.NotNil(t, service_template.New(&netbox.ServiceTemplate{}))
}
