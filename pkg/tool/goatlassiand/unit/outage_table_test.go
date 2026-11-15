package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	atlassianConstant "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/outage_server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnJiraOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.JiraPeer)
		},
		spec,
		webConstant.OutageReaches,
		map[string]webConstant.OutageReach{
			"ListSpaces":      webConstant.OutageElsewhere,
			"SearchPages":     webConstant.OutageElsewhere,
			"CreatePage":      webConstant.OutageElsewhere,
			"GetPage":         webConstant.OutageElsewhere,
			"UpdatePage":      webConstant.OutageElsewhere,
			"DeletePage":      webConstant.OutageElsewhere,
			"GetPageChildren": webConstant.OutageElsewhere,
			"AddPageComment":  webConstant.OutageElsewhere,
			"GetPageDraft":    webConstant.OutageElsewhere,
			"SetPageStatus":   webConstant.OutageElsewhere,
			"EditPage":        webConstant.OutageElsewhere,
			"ListPages":       webConstant.OutageElsewhere,
		},
		nil,
		map[string]string{"UpdateIssue": `{"summary":"alfa"}`},
	)
}

func TestOperationsOnConfluenceOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.ConfluencePeer)
		},
		spec,
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			"ListSpaces":      webConstant.OutageReaches,
			"SearchPages":     webConstant.OutageReaches,
			"CreatePage":      webConstant.OutageReaches,
			"GetPage":         webConstant.OutageReaches,
			"UpdatePage":      webConstant.OutageReaches,
			"DeletePage":      webConstant.OutageReaches,
			"GetPageChildren": webConstant.OutageReaches,
			"AddPageComment":  webConstant.OutageReaches,
			"GetPageDraft":    webConstant.OutageReaches,
			"SetPageStatus":   webConstant.OutageReaches,
			"EditPage":        webConstant.OutageReaches,
			"ListPages":       webConstant.OutageReaches,
		},
		nil,
		map[string]string{"SetPageStatus": `{"status":"current"}`},
	)
}

func TestToolsOnJiraOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.JiraPeer)
		},
		webConstant.OutageReaches,
		map[string]webConstant.OutageReach{
			constant.ConfluenceAddComment:      webConstant.OutageElsewhere,
			constant.ConfluenceCreatePage:      webConstant.OutageElsewhere,
			constant.ConfluenceDeletePage:      webConstant.OutageElsewhere,
			constant.ConfluenceEditPage:        webConstant.OutageElsewhere,
			constant.ConfluenceGetPage:         webConstant.OutageElsewhere,
			constant.ConfluenceGetPageChildren: webConstant.OutageElsewhere,
			constant.ConfluenceGetPageDraft:    webConstant.OutageElsewhere,
			constant.ConfluenceListPages:       webConstant.OutageElsewhere,
			constant.ConfluenceListSpaces:      webConstant.OutageElsewhere,
			constant.ConfluenceSearch:          webConstant.OutageElsewhere,
			constant.ConfluenceSetPageStatus:   webConstant.OutageElsewhere,
			constant.ConfluenceUpdatePage:      webConstant.OutageElsewhere,
		},
		map[string]map[string]any{
			constant.JiraUpdateIssue: {constant.Summary: "alfa"},
		},
	)
}

func TestToolsOnConfluenceOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.ConfluencePeer)
		},
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			constant.ConfluenceAddComment:      webConstant.OutageReaches,
			constant.ConfluenceCreatePage:      webConstant.OutageReaches,
			constant.ConfluenceDeletePage:      webConstant.OutageReaches,
			constant.ConfluenceEditPage:        webConstant.OutageReaches,
			constant.ConfluenceGetPage:         webConstant.OutageReaches,
			constant.ConfluenceGetPageChildren: webConstant.OutageReaches,
			constant.ConfluenceGetPageDraft:    webConstant.OutageReaches,
			constant.ConfluenceListPages:       webConstant.OutageReaches,
			constant.ConfluenceListSpaces:      webConstant.OutageReaches,
			constant.ConfluenceSearch:          webConstant.OutageReaches,
			constant.ConfluenceSetPageStatus:   webConstant.OutageReaches,
			constant.ConfluenceUpdatePage:      webConstant.OutageReaches,
		},
		map[string]map[string]any{
			constant.ConfluenceCreatePage: {
				generativeConstant.ParameterTitle: "alfa",
			},
			constant.ConfluenceSetPageStatus: {
				constant.PageStatus: atlassianConstant.ConfluenceCurrentStatus,
			},
		},
	)
}
