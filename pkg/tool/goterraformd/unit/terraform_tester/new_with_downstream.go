package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	provisionMock "github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"testing"
)

func NewWithDownstream(
	t *testing.T,
	c *mock_client.Client,
	plan string,
	downstream *provisionMock.Client,
) *Tester {
	t.Helper()

	return build(t, c, constant.FixtureHostname, 0, plan, downstream)
}
