package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/system_number"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestNumber(t *testing.T) {
	assert.NotNil(t, system_number.New(&netbox.ASN{}))
}
