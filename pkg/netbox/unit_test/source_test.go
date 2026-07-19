package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/source"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestSource(t *testing.T) {
	assert.NotNil(t, source.New(&netbox.DataSource{}))
}
