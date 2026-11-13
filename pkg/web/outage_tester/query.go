package outage_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/getkin/kin-openapi/openapi3"
	"net/url"
)

func query(
	o *openapi3.Operation,
	parameters map[string]string,
) string {
	result := url.Values{}

	for _, p := range o.Parameters {
		if p.Value == nil || p.Value.In != openapi3.ParameterInQuery ||
			!p.Value.Required {
			continue
		}

		v, okay := parameters[p.Value.Name]

		if !okay {
			v = fmt.Sprint(schemaPlaceholder(p.Value.Schema))
		}

		result.Set(p.Value.Name, v)
	}

	if len(result) == 0 {
		return ""
	}

	return join.Empty("?", result.Encode())
}
