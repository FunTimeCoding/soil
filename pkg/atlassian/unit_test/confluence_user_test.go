package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/user"
	"testing"
)

func TestConfluenceUser(t *testing.T) {
	assert.NotNil(t, user.New(response.NewUser()))
}
