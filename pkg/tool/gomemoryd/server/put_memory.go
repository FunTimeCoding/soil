package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (s *Server) PutMemory(
	_ context.Context,
	r server.PutMemoryRequestObject,
) (server.PutMemoryResponseObject, error) {
	existing, e := s.service.GetMemory(r.Identifier)

	if not_found.Is(e) {
		return server.PutMemory404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PutMemory500JSONResponse(*s.answer(e)), nil
	}

	o := save_option.New()
	o.Name = valueOr(r.Body.Name, existing.Name)
	o.Content = valueOr(r.Body.Content, existing.Content)
	o.Description = valueOr(r.Body.Description, existing.Description)

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

	m, _, f := s.service.UpdateMemory(r.Identifier, o)

	if not_found.Is(f) {
		return server.PutMemory404JSONResponse{Error: f.Error()}, nil
	}

	if conflict.Is(f) {
		return server.PutMemory409JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.PutMemory500JSONResponse(*s.answer(f)), nil
	}

	return server.PutMemory200JSONResponse{Identifier: m.Identifier}, nil
}
