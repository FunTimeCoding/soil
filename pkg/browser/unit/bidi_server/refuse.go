package bidi_server

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/types/message"
)

func (s *Server) Refuse(
	method string,
	code string,
	text string,
) {
	r := message.NewReply()
	r.Type = constant.ErrorType
	r.Error = code
	r.Message = text
	s.mutex.Lock()
	s.answers[method] = r
	s.mutex.Unlock()
}
