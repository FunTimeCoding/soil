package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/service"
	"github.com/funtimecoding/soil/pkg/netbox/service_template"
	"github.com/funtimecoding/soil/pkg/netbox/system_number"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestAddressEntities(t *testing.T) {
	assert.NotNil(
		t,
		internet_address.New(&netbox.IPAddress{Display: "10.0.0.1/24"}),
	)
	assert.NotNil(t, prefix.New(&netbox.Prefix{}))
	assert.NotNil(t, service_template.New(&netbox.ServiceTemplate{}))
	assert.NotNil(t, service.New(&netbox.Service{}))
	assert.NotNil(t, system_number.New(&netbox.ASN{}))
}
