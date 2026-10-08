package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionWindow(
	_ context.Context,
	r server.GetSessionWindowRequestObject,
) (server.GetSessionWindowResponseObject, error) {
	count := 0

	if r.Params.Count != nil {
		count = *r.Params.Count
	}

	blocks, e := s.service.ReadConversation(
		r.Identifier,
		r.Params.Around,
		count,
	)

	if e != nil {
		return server.GetSessionWindow500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	if len(blocks) == 0 {
		return server.GetSessionWindow404JSONResponse(
			server.Error{
				Error: fmt.Sprintf(
					"no block %s in conversation %s",
					r.Params.Around,
					r.Identifier,
				),
			},
		), nil
	}

	result := []server.WindowBlock{}

	for _, b := range blocks {
		result = append(
			result,
			server.WindowBlock{
				Identifier: b.Identifier,
				Role:       b.Role,
				Kind:       b.Kind,
				At:         b.At,
				Text:       b.Text,
			},
		)
	}

	return server.GetSessionWindow200JSONResponse{Blocks: result}, nil
}
