package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *Tester {
	t.Helper()

	return build(
		t,
		c,
		constant.FixtureHostname,
		0,
		constant.FixtureEmptyPlan,
		"",
	)
}
