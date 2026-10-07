package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/types/edit_link_option"
)

func (s *Server) EditLink(
	_ context.Context,
	r server.EditLinkRequestObject,
) (server.EditLinkResponseObject, error) {
	o := edit_link_option.Option{}

	if r.Body.Name != nil {
		o.Name = *r.Body.Name
	}

	if r.Body.Link != nil {
		o.Link = *r.Body.Link
	}

	if r.Body.Description != nil {
		o.Description = *r.Body.Description
	}

	if r.Body.Tags != nil {
		o.Tags = *r.Body.Tags
	}

	if r.Body.Lists != nil {
		o.Lists = *r.Body.Lists
	}

	result, e := s.service.EditLink(int(r.Identifier), o)

	if e != nil {
		return server.EditLink500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.EditLink200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
