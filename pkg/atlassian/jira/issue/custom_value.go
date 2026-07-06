package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/custom_field_value"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (i *Issue) CustomValue(field string) string {
	f := i.fieldMap.ByName(field)

	if f == nil {
		return UnknownField
	}

	verbose := i.option.Verbose

	for k, v := range i.Raw.Fields.Unknowns {
		if k == f.ID {
			if v == nil {
				return NilValue
			}

			switch cast := v.(type) {
			case float64:
				if verbose {
					fmt.Println("Float value")
				}

				return fmt.Sprintf("%v", v)
			case string:
				if verbose {
					fmt.Println("String value")
				}

				return cast
			case map[string]any:
				if verbose {
					fmt.Println("String-any map value")
				}

				return custom_field_value.FromMap(cast).Value
			case []any:
				if verbose {
					fmt.Println("Any slice value")
				}

				var result []string

				for _, item := range v.([]any) {
					switch castInner := item.(type) {
					case map[string]any:
						if verbose {
							fmt.Println("Map value")
						}

						result = append(
							result,
							custom_field_value.FromMap(castInner).Value,
						)
					default:
						fmt.Printf("Unexpected type inner: %T\n", item)
					}
				}

				strings.Sort(result, true)

				return join.Comma(result)
			default:
				fmt.Printf("Unexpected type: %T\n", v)

				return fmt.Sprintf("%+v", v)
			}
		}
	}

	return UnknownValue
}
