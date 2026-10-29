package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"testing"
)

func NewAs(
	t *testing.T,
	c *mock_client.Client,
	hostname string,
) *Tester {
	t.Helper()

	return build(t, c, hostname, 0, constant.FixtureEmptyPlan, "")
}
