package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim/message"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"github.com/funtimecoding/soil/pkg/system"
	gomonitord "github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	"testing"
)

func TestWindowReleasesItsClaimsOnItemsGone(t *testing.T) {
	s := base.New(t)
	connectTo(t, s.Port)
	m := monitor.New(true, memory.New())
	defer m.Close()
	owner := fmt.Sprintf(
		monitorConstant.OwnerFormat,
		system.User().Username,
		system.Hostname(),
	)
	d := gomonitord.NewEnvironment()
	defer d.Close()
	x := t.Context()
	_, e := d.PutClaim(x, "jira-GONE-1", owner)
	assert.FatalOnError(t, e)
	_, f := d.PutClaim(x, "jira-GONE-1", "someone@elsewhere")
	assert.FatalOnError(t, f)
	m.Update(
		message.Message{
			Claims: []client.Claim{
				{Item: "jira-GONE-1", Owner: owner},
				{Item: "jira-GONE-1", Owner: "someone@elsewhere"},
			},
		},
	)
	present := item.New(
		monitorConstant.GoJira,
		"jira-HERE-1",
		soil.Warning,
		"still here",
		"",
		nil,
	)
	present.Label = present.Identifier
	_, c := m.Update(fetch.Message{Items: []*item.Item{present}})
	runCommand(c)
	r, g := d.GetClaims(x)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, r)
	assert.String(t, "someone@elsewhere", r[0].Owner)
}
