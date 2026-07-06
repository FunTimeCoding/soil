package user

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"testing"
)

func TestUser(t *testing.T) {
	assert.NotNil(t, New(response.NewUser()))
}
