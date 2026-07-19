package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestSite(t *testing.T) {
	assert.NotNil(t, site.New(&netbox.Site{}))
}
