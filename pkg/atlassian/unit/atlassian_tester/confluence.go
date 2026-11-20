package atlassian_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (o *Tester) Confluence(server string) *basic.Client {
	o.t.Helper()
	result := basic.New(locator.NewOrigin(server), "alfa", "bravo-token")
	o.t.Cleanup(result.Close)

	return result
}
