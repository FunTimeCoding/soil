package outage_tester

import "github.com/getkin/kin-openapi/openapi3"

func schemaPlaceholder(s *openapi3.SchemaRef) any {
	if s == nil || s.Value == nil {
		return "1"
	}

	p := property{Enum: s.Value.Enum}

	if s.Value.Type != nil && len(*s.Value.Type) > 0 {
		p.Type = (*s.Value.Type)[0]
	}

	return placeholder(p)
}
