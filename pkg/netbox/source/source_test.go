package source

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestSource(t *testing.T) {
	assert.NotNil(t, New(&netbox.DataSource{}))
}
