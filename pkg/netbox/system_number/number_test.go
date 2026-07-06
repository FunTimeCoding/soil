package system_number

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestNumber(t *testing.T) {
	assert.NotNil(t, New(&netbox.ASN{}))
}
