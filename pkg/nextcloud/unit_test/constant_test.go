package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/nextcloud/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "NEXTCLOUD_HOST", constant.HostEnvironment)
	assert.String(t, "NEXTCLOUD_USER", constant.UserEnvironment)
	assert.String(t, "NEXTCLOUD_PASSWORD", constant.PasswordEnvironment)
}
