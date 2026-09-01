package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/custom_field_value"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (i *Issue) CustomValue(field string) string {
	f := i.fieldMap.ByName(field)

	if f == nil {
		return constant.JiraUnknownField
	}

	verbose := i.option.Verbose

	for k, v := range i.Raw.Fields.Unknowns {
		if k == f.ID {
			if v == nil {
				return constant.JiraNilValue
			}

			switch cast := v.(type) {
			case float64:
				if verbose {
					console.Line("Float value")
				}

				return fmt.Sprintf("%v", v)
			case string:
				if verbose {
					console.Line("String value")
				}

				return cast
			case map[string]any:
				if verbose {
					console.Line("String-any map value")
				}

				return custom_field_value.FromMap(cast).Value
			case []any:
				if verbose {
					console.Line("Any slice value")
				}

				var result []string

				for _, item := range v.([]any) {
					switch castInner := item.(type) {
					case map[string]any:
						if verbose {
							console.Line("Map value")
						}

						result = append(
							result,
							custom_field_value.FromMap(castInner).Value,
						)
					default:
						console.Format("Unexpected type inner: %T\n", item)
					}
				}

				strings.Sort(result, true)

				return join.Comma(result)
			default:
				console.Format("Unexpected type: %T\n", v)

				return fmt.Sprintf("%+v", v)
			}
		}
	}

	return constant.JiraUnknownValue
}
