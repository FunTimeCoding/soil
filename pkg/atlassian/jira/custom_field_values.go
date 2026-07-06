package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/custom_field_value"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) CustomFieldValues(
	key string,
	issueType string,
	fieldName string,
) ([]custom_field_value.Value, error) {
	m, e := c.FieldMap()

	if e != nil {
		return nil, e
	}

	f := m.ByName(fieldName)

	if f == nil {
		return nil, fmt.Errorf(
			"field not found: %s: %w",
			fieldName,
			constant.ErrorNotFound,
		)
	}

	meta, g := c.CreateMeta(key)

	if g != nil {
		return nil, g
	}

	for _, p := range meta.Projects {
		for _, t := range p.IssueTypes {
			if t.Name != issueType {
				continue
			}

			var result []custom_field_value.Value
			notation.MustDecode(
				notation.Encode(
					t.Fields[f.ID].(map[string]any)[constant.AllowedValuesKey],
					false,
				),
				&result,
				true,
			)

			return result, nil
		}
	}

	return nil, fmt.Errorf(
		"issue type not found: %s: %w",
		issueType,
		constant.ErrorNotFound,
	)
}
