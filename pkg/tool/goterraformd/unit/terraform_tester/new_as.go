package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/mock_client"
	"testing"
)

func NewAs(
	t *testing.T,
	c *mock_client.Client,
	hostname string,
) *Tester {
	t.Helper()

	return build(t, c, hostname, 0)
}
