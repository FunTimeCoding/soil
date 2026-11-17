package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/client"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestClientReachesGomaintlogd(t *testing.T) {
	assert.FatalOnError(
		t,
		base.New(t).MaintenanceClient().Reachable(t.Context()),
	)
}

func TestRefusedTokenIsNotReachable(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port),
		"wrong-token",
	)
	t.Cleanup(c.Close)
	assert.Error(t, c.Reachable(t.Context()))
}
