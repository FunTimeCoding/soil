package outage_tester

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/getkin/kin-openapi/openapi3"
)

func body(o *openapi3.Operation) string {
	if o.RequestBody == nil || o.RequestBody.Value == nil {
		return "{}"
	}

	media := o.RequestBody.Value.Content.Get(constant.Object)

	if media == nil || media.Schema == nil || media.Schema.Value == nil {
		return "{}"
	}

	result := map[string]any{}

	for _, name := range media.Schema.Value.Required {
		result[name] = schemaPlaceholder(media.Schema.Value.Properties[name])
	}

	return string(notation.Marshal(result))
}
