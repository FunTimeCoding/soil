package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user"
	rawUser "github.com/opsgenie/opsgenie-go-sdk-v2/user"
	"testing"
)

func TestOpsgenieUser(t *testing.T) {
	assert.NotNil(t, user.New(&rawUser.User{}))
}
