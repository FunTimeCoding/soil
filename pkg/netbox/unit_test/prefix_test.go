package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestPrefix(t *testing.T) {
	assert.NotNil(t, prefix.New(&netbox.Prefix{}))
}
