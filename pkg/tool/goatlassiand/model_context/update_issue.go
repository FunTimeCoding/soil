package model_context

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/trivago/tgo/tcontainer"
)

func (s *Server) updateIssue(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	summary := r.GetString(constant.Summary, "")
	description := r.GetString("description", "")
	assigneeName := r.GetString(constant.Assignee, "")
	reporterName := r.GetString(constant.Reporter, "")
	labelsRaw := r.GetString(constant.Labels, "")
	fieldsRaw := r.GetString(constant.AdditionalFields, "")
	noDiff := r.GetBool(constant.NoDiff, false)

	if summary == "" &&
		description == "" &&
		assigneeName == "" &&
		reporterName == "" &&
		labelsRaw == "" &&
		fieldsRaw == "" {
		return response.Fail("no fields to update")
	}

	before, g := s.jira.Issue(key)

	if g != nil {
		return s.captureFail(g, "issue not found")
	}

	raw := issue.Raw(key)
	raw.Fields.Unknowns = make(tcontainer.MarshalMap)

	if summary != "" {
		raw.Fields.Summary = summary
	}

	if description != "" {
		raw.Fields.Description = description
	}

	if reporterName != "" {
		user, fail, g := s.resolveAssignee(c, reporterName)

		if fail != nil {
			return fail, g
		}

		raw.Fields.Reporter = user
	}

	if labelsRaw != "" {
		var labels []string

		if g := json.Unmarshal(
			[]byte(labelsRaw),
			&labels,
		); g != nil {
			return response.Fail(
				"labels must be a JSON array of strings: %v",
				g,
			)
		}

		raw.Fields.Labels = labels
	}

	var customFieldNames []string

	if fieldsRaw != "" {
		var fields map[string]any

		if g := json.Unmarshal(
			[]byte(fieldsRaw),
			&fields,
		); g != nil {
			return response.Fail(
				"additional_fields must be a JSON object: %v",
				g,
			)
		}

		fieldMap, h := s.jira.FieldMap()

		if h != nil {
			return s.captureDetail(h)
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
			customFieldNames = append(customFieldNames, name)
		}
	}

	hasFieldChanges := summary != "" ||
		description != "" ||
		labelsRaw != "" ||
		fieldsRaw != ""

	if hasFieldChanges {
		_, resp, i := s.jira.Nested().Issue.UpdateWithContext(
			c,
			raw,
		)

		if i != nil {
			if resp != nil && resp.Body != nil {
				return s.captureFail(
					i,
					fmt.Sprintf(
						"update failed: %s",
						string(system.ReadAll(resp.Body)),
					),
				)
			}

			return s.captureFail(i, "issue not updated")
		}
	}

	if assigneeName != "" {
		user, fail, g := s.resolveAssignee(c, assigneeName)

		if fail != nil {
			return fail, g
		}

		resp, j := s.jira.Nested().Issue.UpdateAssigneeWithContext(
			c,
			key,
			user,
		)

		if j != nil {
			if resp != nil && resp.Body != nil {
				return s.captureFail(
					j,
					fmt.Sprintf(
						"assign failed: %s",
						string(system.ReadAll(resp.Body)),
					),
				)
			}

			return s.captureFail(j, "assignment not applied")
		}
	}

	after, k := s.jira.Issue(key)

	if k != nil {
		return s.captureFail(k, "issue updated but retrieval failed")
	}

	return response.SuccessAny(
		convert.JiraIssueDiff(before, after, noDiff, customFieldNames),
	)
}
