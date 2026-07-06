package console_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	port int,
) *Tester {
	t.Helper()
	c, e := client.NewClientWithResponses(
		locator.New(
			constant.Localhost,
		).Insecure().Port(port).String(),
	)
	assert.FatalOnError(t, e)

	return &Tester{t: t, client: c}
}
