package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestClusterType(t *testing.T) {
	assert.NotNil(t, cluster_type.New(&netbox.ClusterType{}))
}
