package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestClientReachesDaemon(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestProbeRefusesUnknownToken(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		"bogus",
	)
	defer c.Close()
	assert.NotNil(t, c.Reachable(t.Context()))
}

func TestClaimWithoutOwnerIsValidation(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	_, e := c.PutClaim(t.Context(), "jira-ABC-1", "")
	assert.True(t, validation.Is(e))
	assert.True(t, validation.Is(c.DeleteClaim(t.Context(), "jira-ABC-1", "")))
}
