package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "SENTRY_LOCATOR", constant.LocatorEnvironment)
	assert.String(t, "SENTRY_ORGANIZATION", constant.OrganizationEnvironment)
	assert.String(t, "SENTRY_PROJECT", constant.ProjectEnvironment)
}
