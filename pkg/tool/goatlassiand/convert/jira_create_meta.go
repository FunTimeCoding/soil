package convert

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/custom_field_value"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"sort"
)

func JiraCreateMeta(
	t *jira.MetaIssueType,
	expand []string,
) []*CreateMetaField {
	expandSet := make(map[string]bool, len(expand))

	for _, n := range expand {
		expandSet[n] = true
	}

	var result []*CreateMetaField

	for key := range t.Fields {
		f := &CreateMetaField{Key: key}
		name, e := t.Fields.String(join.Empty(key, "/name"))

		if e == nil {
			f.Name = name
		}

		required, g := t.Fields.Bool(join.Empty(key, "/required"))

		if g == nil {
			f.Required = required
		}

		schema, h := t.Fields.String(join.Empty(key, "/schema/type"))

		if h == nil {
			f.Schema = schema
		}

		allowed, i := t.Fields.Array(
			join.Empty(key, "/allowedValues"),
		)

		if i == nil {
			for _, a := range allowed {
				m, okay := a.(map[string]any)

				if !okay {
					continue
				}

				var v custom_field_value.Value
				notation.MustDecode(
					notation.Encode(m, false),
					&v,
					true,
				)

				if v.Value != "" {
					f.AllowedValues = append(
						f.AllowedValues,
						&CreateMetaAllowed{
							Identifier: v.ID,
							Value:      v.Value,
						},
					)
				} else if n, okay := m["name"].(string); okay {
					f.AllowedValues = append(
						f.AllowedValues,
						&CreateMetaAllowed{
							Identifier: v.ID,
							Value:      n,
						},
					)
				}
			}
		}

		if len(f.AllowedValues) > constant.AllowedValueLimit &&
			!expandSet[f.Name] && !expandSet[f.Key] {
			sort.Slice(
				f.AllowedValues,
				func(a, b int) bool {
					return strings.MustToInteger(f.AllowedValues[a].Identifier) >
						strings.MustToInteger(f.AllowedValues[b].Identifier)
				},
			)
			f.AllowedValues = f.AllowedValues[:constant.AllowedValueLimit]
		}

		result = append(result, f)
	}

	sort.Slice(
		result,
		func(a, b int) bool {
			if result[a].Required != result[b].Required {
				return result[a].Required
			}

			return result[a].Name < result[b].Name
		},
	)

	return result
}
