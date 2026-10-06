package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostChunk(
	_ context.Context,
	r server.PostChunkRequestObject,
) (server.PostChunkResponseObject, error) {
	p := s.service.Preview(r.Body.Path, r.Body.Body)
	chunks := make([]server.PreviewChunk, len(p.Chunks))

	for i, c := range p.Chunks {
		chunks[i] = server.PreviewChunk{
			Index:        c.Index,
			FirstLine:    c.FirstLine,
			LastLine:     c.LastLine,
			Bytes:        c.Bytes,
			Tokens:       c.Tokens,
			Piece:        c.Piece,
			CutChecked:   c.CutChecked,
			CutLevel:     c.CutLevel,
			CutLines:     lines(c.CutLines),
		}
	}

	sections := make([]server.PreviewSection, len(p.Sections))

	for i, x := range p.Sections {
		sections[i] = convertPreviewSection(x)
	}

	return server.PostChunk200JSONResponse{
		Model:     p.Model,
		Window:    p.Window,
		Allowance: p.Allowance,
		Chunks:    chunks,
		Sections:  sections,
	}, nil
}
