package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) AddLink(
	_ context.Context,
	r server.AddLinkRequestObject,
) (server.AddLinkResponseObject, error) {
	value := ""

	if r.Body.List != nil {
		value = *r.Body.List
	}

	listIdentifier, e := s.service.ResolveList(value)

	if e != nil {
		return server.AddLink500JSONResponse(*s.captureDetail(e)), nil
	}

	name := ""

	if r.Body.Name != nil {
		name = *r.Body.Name
	}

	if name == "" {
		name = r.Body.Link
	}

	result, f := s.client.CreateLink(r.Body.Link, name, listIdentifier, nil)

	if f != nil {
		return server.AddLink500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.AddLink200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
