package integration

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/integration/base"
	"testing"
)

func TestAWindowReleasesItsClaimsOnItemsGone(t *testing.T) {
	s := base.New(t)
	connectTo(t, s.Port)
	m := monitor.New(true)
	owner := fmt.Sprintf(
		monitorConstant.OwnerFormat,
		system.User().Username,
		system.Hostname(),
	)
	_, e := s.Store.ClaimItem("jira-GONE-1", owner)
	assert.FatalOnError(t, e)
	_, e = s.Store.ClaimItem("jira-GONE-1", "someone@elsewhere")
	assert.FatalOnError(t, e)
	m.Update(
		claim.Message{
			Claims: []client.Claim{
				{Item: "jira-GONE-1", Owner: owner},
				{Item: "jira-GONE-1", Owner: "someone@elsewhere"},
			},
		},
	)
	present := item.New(
		monitorConstant.GoJira,
		"jira-HERE-1",
		constant.Warning,
		"still here",
		"",
		nil,
	)
	present.Label = present.Identifier
	_, c := m.Update(fetch.Message{Items: []*item.Item{present}})
	runCommand(c)
	v, e := s.Store.Claims()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "someone@elsewhere", v[0].Owner)
}
