package model_context

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/trivago/tgo/tcontainer"
)

func (s *Server) createIssue(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	project, f := r.RequireString(constant.Project)

	if f != nil {
		return response.Fail("project is required: %v", f)
	}

	issueType, g := r.RequireString(constant.IssueType)

	if g != nil {
		return response.Fail("issue_type is required: %v", g)
	}

	summary, h := r.RequireString(constant.Summary)

	if h != nil {
		return response.Fail("summary is required: %v", h)
	}

	description := r.GetString("description", "")
	assigneeName := r.GetString(constant.Assignee, "")
	labelsRaw := r.GetString(constant.Labels, "")
	fieldsRaw := r.GetString(constant.AdditionalFields, "")
	raw := issue.RawStub()
	raw.Fields.Unknowns = make(tcontainer.MarshalMap)
	reporter, i := s.jira.User()

	if i != nil {
		return s.captureDetail(i)
	}

	raw.Fields.Reporter = reporter
	raw.Fields.Project.Key = project
	raw.Fields.Type.Name = issueType
	raw.Fields.Summary = summary
	raw.Fields.Description = description

	if labelsRaw != "" {
		var labels []string

		if i := json.Unmarshal(
			[]byte(labelsRaw),
			&labels,
		); i != nil {
			return response.Fail(
				"labels must be a JSON array of strings: %v",
				i,
			)
		}

		raw.Fields.Labels = labels
	}

	if fieldsRaw != "" {
		var fields map[string]any

		if i := json.Unmarshal(
			[]byte(fieldsRaw),
			&fields,
		); i != nil {
			return response.Fail(
				"additional_fields must be a JSON object: %v",
				i,
			)
		}

		fieldMap, j := s.jira.FieldMap()

		if j != nil {
			return s.captureDetail(j)
		}

		for name, value := range fields {
			field := fieldMap.ByName(name)

			if field == nil {
				return response.Fail(
					"unknown field: %s",
					name,
				)
			}

			raw.Fields.Unknowns.Set(field.Key, value)
		}
	}

	result, resp, k := s.jira.Nested().Issue.CreateWithContext(
		c,
		raw,
	)

	if k != nil {
		if resp != nil && resp.Body != nil {
			return s.captureFail(
				k,
				fmt.Sprintf(
					"create failed: %s",
					string(system.ReadAll(resp.Body)),
				),
			)
		}

		return s.captureFail(k, "issue not created")
	}

	if assigneeName != "" {
		user, fail, j := s.resolveAssignee(c, assigneeName)

		if fail != nil {
			return fail, j
		}

		assignResp, l := s.jira.Nested().Issue.UpdateAssigneeWithContext(
			c,
			result.Key,
			user,
		)

		if l != nil {
			if assignResp != nil && assignResp.Body != nil {
				return s.captureFail(
					l,
					fmt.Sprintf(
						"created %s but assign failed: %s",
						result.Key,
						string(system.ReadAll(assignResp.Body)),
					),
				)
			}

			return s.captureFail(
				l,
				fmt.Sprintf(
					"created %s but assign failed",
					result.Key,
				),
			)
		}
	}

	created, m := s.jira.Issue(result.Key)

	if m != nil {
		return s.captureFail(m, "issue created but retrieval failed")
	}

	return response.SuccessAny(convert.JiraIssue(created))
}
