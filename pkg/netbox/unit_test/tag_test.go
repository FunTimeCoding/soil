package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestTag(t *testing.T) {
	assert.NotNil(t, tag.New(&netbox.Tag{}))
}
