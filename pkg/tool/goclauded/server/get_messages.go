package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
)

func (s *Server) GetMessages(
	x context.Context,
	r server.GetMessagesRequestObject,
) (server.GetMessagesResponseObject, error) {
	var identifiers []uint

	for _, i := range r.Params.Identifier {
		identifiers = append(identifiers, uint(i))
	}

	found, e := s.message.ReadMessages(identifiers)

	if e != nil {
		return server.GetMessages500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	result := server.MessageReadResponse{
		Messages: []server.MessageEntry{},
		Missing:  []int{},
	}

	for _, m := range found {
		result.Messages = append(result.Messages, *messageEntry(m))
	}

	for _, i := range message.Missing(identifiers, found) {
		result.Missing = append(result.Missing, int(i))
	}

	return server.GetMessages200JSONResponse(result), nil
}
