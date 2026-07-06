package user_map

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user"
	"testing"
)

func TestMap(t *testing.T) {
	assert.NotNil(t, New([]*user.User{}))
}
