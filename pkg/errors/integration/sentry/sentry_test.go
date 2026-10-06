//go:build local

package sentry

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func TestWhoamiAndOrganizationsRead(t *testing.T) {
	c := sentry.NewEnvironment()
	u, e := c.Whoami()
	assert.FatalOnError(t, e)
	assert.True(t, u.Identifier != "")
	organizations, f := c.Organizations()
	assert.FatalOnError(t, f)
	assert.NotEmpty(t, organizations)
}

func TestMissingIssueIsNotFound(t *testing.T) {
	_, e := sentry.NewEnvironment().IssueByIdentifier(
		environment.Required(constant.OrganizationEnvironment),
		"999999999",
	)
	assert.True(t, not_found.Is(e))
}
