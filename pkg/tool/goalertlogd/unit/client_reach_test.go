package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestTheClientReachesGoalertlogd(t *testing.T) {
	s := base.New(t)
	assert.FatalOnError(t, s.AlertClient(t).Reachable(t.Context()))
}

func TestARefusedTokenIsNotReachable(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port),
		"wrong-token",
	)
	t.Cleanup(c.Close)
	assert.True(t, c.Reachable(t.Context()) != nil)
}
