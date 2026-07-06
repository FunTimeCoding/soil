package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/trivago/tgo/tcontainer"
)

func (s *Server) writeChecklist(
	c context.Context,
	key string,
	items []*checklist_item.Item,
) (*mcp.CallToolResult, error) {
	m, e := s.jira.FieldMap()

	if e != nil {
		return s.captureDetail(e)
	}

	field := m.ByName(constant.ChecklistField)

	if field == nil {
		return response.Fail("checklist field not found")
	}

	raw := issue.Raw(key)
	raw.Fields.Unknowns = make(tcontainer.MarshalMap)
	raw.Fields.Unknowns.Set(field.Key, formatChecklist(items))
	_, resp, f := s.jira.Nested().Issue.UpdateWithContext(c, raw)

	if f != nil {
		if resp != nil && resp.Body != nil {
			return s.captureFail(
				f,
				fmt.Sprintf(
					"checklist update failed: %s",
					string(system.ReadAll(resp.Body)),
				),
			)
		}

		return s.captureFail(f, "checklist not updated")
	}

	return nil, nil
}
