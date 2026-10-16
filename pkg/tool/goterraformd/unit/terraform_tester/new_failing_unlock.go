package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/mock_client"
	"testing"
)

func NewFailingUnlock(
	t *testing.T,
	c *mock_client.Client,
) *Tester {
	t.Helper()

	return build(
		t,
		c,
		constant.FixtureHostname,
		1,
		constant.FixtureEmptyPlan,
		"",
	)
}
