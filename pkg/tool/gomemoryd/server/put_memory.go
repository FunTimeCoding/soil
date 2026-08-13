package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (s *Server) PutMemory(
	_ context.Context,
	r server.PutMemoryRequestObject,
) (server.PutMemoryResponseObject, error) {
	o := save_option.New()
	o.Name = r.Body.Name
	o.Content = r.Body.Content
	o.Description = r.Body.Description

	if r.Body.Tags != nil {
		o.Tags = *r.Body.Tags
	}

	if r.Body.Metadata != nil {
		o.Metadata = *r.Body.Metadata
	}

	if r.Body.Source != nil {
		o.Source = *r.Body.Source
	}

	if r.Body.ProvenanceHash != nil {
		o.ProvenanceHash = *r.Body.ProvenanceHash
	}

	if r.Body.Ordinal != nil {
		o.Ordinal = *r.Body.Ordinal
	}

	m, e := s.service.UpdateMemory(r.Identifier, o)

	if e != nil {
		return server.PutMemory500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PutMemory200JSONResponse{Identifier: m.Identifier}, nil
}
