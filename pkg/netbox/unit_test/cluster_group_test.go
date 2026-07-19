package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestClusterGroup(t *testing.T) {
	assert.NotNil(t, cluster_group.New(&netbox.ClusterGroup{}))
}
