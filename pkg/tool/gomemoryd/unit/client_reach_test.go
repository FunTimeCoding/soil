package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestTheClientReachesGomemoryd(t *testing.T) {
	s := base.New(t)
	assert.FatalOnError(t, s.MemoryClient(t).Reachable(t.Context()))
}

func TestARefusedTokenIsNotReachable(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port),
		"wrong-token",
	)
	t.Cleanup(c.Close)
	assert.Error(t, c.Reachable(t.Context()))
}
