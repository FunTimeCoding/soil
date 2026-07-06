package user

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/user"
	"testing"
)

func TestUser(t *testing.T) {
	assert.NotNil(t, New(&user.User{}))
}
