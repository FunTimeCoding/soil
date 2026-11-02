package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/item"
	"github.com/funtimecoding/soil/pkg/system"
	gomonitord "github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestAWindowReleasesItsClaimsOnItemsGone(t *testing.T) {
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
	assert.FatalOnError(t, d.ClaimItem("jira-GONE-1", owner))
	assert.FatalOnError(t, d.ClaimItem("jira-GONE-1", "someone@elsewhere"))
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
		soil.Warning,
		"still here",
		"",
		nil,
	)
	present.Label = present.Identifier
	_, c := m.Update(fetch.Message{Items: []*item.Item{present}})
	runCommand(c)
	g, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)
	r, f := g.GetClaimsWithResponse(t.Context())
	assert.FatalOnError(t, f)
	assert.FatalNotNil(t, r.JSON200)
	assert.Count(t, 1, *r.JSON200)
	assert.String(t, "someone@elsewhere", (*r.JSON200)[0].Owner)
}
