package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user_map"
	"testing"
)

func TestOpsgenieUserMap(t *testing.T) {
	assert.NotNil(t, user_map.New([]*user.User{}))
}
