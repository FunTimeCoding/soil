package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"strings"
)

func (s *Server) GetCreateMeta(
	_ context.Context,
	r server.GetCreateMetaRequestObject,
) (server.GetCreateMetaResponseObject, error) {
	var expand []string

	if r.Params.Expand != nil {
		for _, v := range strings.Split(*r.Params.Expand, ",") {
			expand = append(expand, strings.TrimSpace(v))
		}
	}

	t, e := s.service.CreateMeta(r.Params.Project, r.Params.IssueType)

	if e != nil {
		return server.GetCreateMeta500JSONResponse(*s.captureDetail(e)), nil
	}

	var result server.GetCreateMeta200JSONResponse

	for _, f := range convert.JiraCreateMeta(t, expand) {
		field := &server.CreateMetaField{
			Name:     f.Name,
			Key:      f.Key,
			Required: f.Required,
			Schema:   f.Schema,
		}

		if len(f.AllowedValues) > 0 {
			var allowed []server.CreateMetaAllowed

			for _, a := range f.AllowedValues {
				allowed = append(
					allowed,
					server.CreateMetaAllowed{
						Identifier: a.Identifier,
						Value:      a.Value,
					},
				)
			}

			field.AllowedValues = &allowed
		}

		result = append(result, field)
	}

	return result, nil
}
